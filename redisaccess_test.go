package redisaccess

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	_, err := InitRedisLocalhost()
	if err == nil {
		list, err := GetKeys("share-files::*")
		if err == nil {
			for _, key := range list {
				fmt.Println(key)
			}
		}
	} else {
		fmt.Println("Error in InitRedis: ", err.Error())
	}

}
