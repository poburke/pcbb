package shared

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "redis:6379", // Address of Redis container
        Password: "",            // no password set
        DB:       0,             // use default DB
    })
    return rdb
}
