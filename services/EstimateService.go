package services

import (
	"arman-estimation-service/types/structs"
	. "arman-estimation-service/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"time"
)

type EstimateService struct {
	RedisClient *redis.Client
}

func (r EstimateService) Estimate(ctx context.Context, segment string) (uint64, *structs.CustomError) {

	// At first if segment exists then return the value
	_, err := r.RedisClient.Get(ctx, segment).Result()
	if err == nil {
		// just count elements of already created key.
		// This prevents redundant PFMERGE
		return r.doPFCount(ctx, segment)
	}

	// in case of any error return to higher stack frame
	if err != redis.Nil {
		return 0, CustomError(structs.Categories.Internal, err.Error())
	}

	// find keys of current segment for N days ago. Keys would be an array of strings (segment:date)
	// {"sport:20220912", "sport:20220911", "sport:20220910", "sport:20220909", ...}
	keys := r.findKeys(segment)

	// This line use redis `pfmerge` command to aggregate distinct values of previous created keys together.
	// the final command would be something like this:
	// PFMERGE sport sport:20220912 sport:20220911 sport:20220910 sport:20220909 ...
	err = r.RedisClient.PFMerge(ctx, segment, keys...).Err()
	if err != nil {
		return 0, CustomError(structs.Categories.Internal, err.Error())
	}

	// Set expiration for created segment. This will prevent too many "PFMERGE"s overhead.
	// It will cache last 14 days counters into a new key.
	err = r.RedisClient.Expire(ctx, segment, time.Hour).Err()
	if err != nil {
		return 0, CustomError(structs.Categories.Internal, err.Error())
	}

	// Finally count elements of new created key
	return r.doPFCount(ctx, segment)
}

func (r EstimateService) findKeys(segment string) []string {
	today := time.Now()
	days := viper.GetInt("KEEP_USER_SEGMENT_DAYS")

	// iterate of last N days (current is 14).
	// I don't use "KEY sport:*" command because of its poor performance
	var result []string
	for i := 0; i < days; i++ {
		key := fmt.Sprintf(
			"%s:%s",
			segment,
			today.Add(time.Duration(-i)*time.Hour*24).Format("20060102"),
		)
		result = append(result, key)
	}

	return result
}

func (r EstimateService) doPFCount(ctx context.Context, segment string) (uint64, *structs.CustomError) {
	count, err := r.RedisClient.PFCount(ctx, segment).Uint64()
	if err != nil {
		return 0, CustomError(structs.Categories.Internal, err.Error())
	}

	return count, nil
}
