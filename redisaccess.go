// redisaccess project redisaccess.go
package redisaccess

import (
	"encoding/json"
	"errors"
	"strings"

	"time"

	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func isInitialized() (err error) {
	if redisClient == nil {
		err = errors.New("Redis not initialized, call InitRedis first")
	}
	return
}

func InitRedis(servername, password string) (client *redis.Client) {

	if !strings.Contains(servername, ":") {
		servername += ":6379"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     servername,
		Password: password,
		DB:       0, // use default DB
	})
	client = redisClient
	return

}

func SetValue(key string, value interface{}, duration time.Duration) (err error) {
	err = isInitialized()
	if err == nil {
		data, err := json.Marshal(value)
		if err == nil {

			status := redisClient.Set(key, data, duration)

			err = status.Err()
		}
	}

	return
}

func GetValue(key string) (value string, found bool, err error) {
	err = isInitialized()
	if err == nil {
		result := redisClient.Get(key)

		err = result.Err()
		found = err == nil
		if found {

			value = result.Val()
		}
	}
	return
}

func RemoveValue(key string) (err error) {
	err = isInitialized()
	if err == nil {
		status := redisClient.Del(key)
		err = status.Err()
	}

	return
}

func AddToQueue(queuename string, key string, value interface{}) (success bool, err error) {
	err = isInitialized()
	if err == nil {
		data, err := json.Marshal(value)
		success = err == nil
		if success {

			cmd := redisClient.HSet(queuename, key, data)
			err = cmd.Err()
			success = cmd.Err() == nil
		}
	}
	return
}

func ReadQueue(queuename string) (queue []string, err error) {
	err = isInitialized()
	if err == nil {
		queue, err = redisClient.HKeys(queuename).Result()
	}

	return
}

func RemoveFromQueue(queuename, key string) (err error) {
	err = isInitialized()
	if err == nil {
		cmd := redisClient.HDel(queuename, key)
		err = cmd.Err()
	}
	return
}

func ScanQueue(queuename string, limit int) (queue []string, err error) {
	err = isInitialized()
	if err == nil {
		queue, _, err = redisClient.HScan(queuename, 0, "", int64(limit)).Result()
	}

	return
}
