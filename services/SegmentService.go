package services

import (
	"arman-estimation-service/types/structs"
	. "arman-estimation-service/utils"
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"time"
)

type SegmentService struct {
	RedisClient *redis.Client
}

func (r *SegmentService) Push(ctx context.Context, segment string, userId uint32) *structs.CustomError {

	// Make a key by combination of segment name and current date
	date := time.Now().Format("20060102")
	key := segment + ":" + date

	// Use "get" instead of "exists" command, because "Get" time complexity is O(1)
	_, err := r.RedisClient.Get(ctx, key).Result()

	switch err {
	case redis.Nil:
		// In the first time creation of this key, we add current user with expiration of n days.
		// Number of days will be read from env val "KEEP_USER_SEGMENT_DAYS"
		days := viper.GetInt("KEEP_USER_SEGMENT_DAYS")
		return r.addWithTTL(ctx, key, userId, time.Duration(days)*time.Hour*24)
	case nil:
		// For the second time and others we just add this user to the key
		return r.add(ctx, key, userId)
	default:
		// There is an error, so we decorate and return the error to the above stack frame
		return CustomError(structs.Categories.Internal, err.Error())
	}
}

//add current user id into the pair of (segment:current date)
func (r *SegmentService) add(ctx context.Context, key string, userId uint32) *structs.CustomError {
	err := r.RedisClient.PFAdd(ctx, key, userId).Err()
	if err != nil {
		return CustomError(structs.Categories.Internal, err.Error())
	}

	return nil
}

//addWithTTL add current user id into the pair of (segment:current date).
//This will be called when this is first user of this segment in current day.
//So it will set expiration of n days (current is 14 days)
func (r *SegmentService) addWithTTL(ctx context.Context, key string, userId uint32, ttl time.Duration) *structs.CustomError {

	addErr := r.add(ctx, key, userId)
	if addErr != nil {
		return addErr
	}

	expErr := r.RedisClient.Expire(ctx, key, ttl).Err()
	if expErr != nil {
		return CustomError(structs.Categories.Internal, expErr.Error())
	}

	return nil
}
