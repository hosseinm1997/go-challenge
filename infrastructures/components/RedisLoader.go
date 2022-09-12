package components

//InitializeRedisServer initialize redis client connection to redis server only once. (shared with requests)
//It prevents system from opening and closing connection too many times.
import (
	"arman-estimation-service/types/aliases"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"sync"
)

var (
	rdb       *redis.Client
	redisOnce sync.Once
)

func InitializeRedisServer(config aliases.StringMap) {
	if rdb == nil {
		redisOnce.Do(func() {
			rdb = redis.NewClient(&redis.Options{
				Addr:     viper.GetString("REDIS_HOST") + ":" + viper.GetString("REDIS_PORT"),
				Password: viper.GetString("REDIS_PASSWORD"), // no password set
				DB:       viper.GetInt("REDIS_DB"),          // use default DB
			})
		})
	}
}

func GetRedisClient() *redis.Client {
	return rdb
}
