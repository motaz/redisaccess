package redisaccess

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	_, err := InitRedisLocalhost()
	if err == nil {
		err := SetValue("sec-key", "My Data-3", time.Second*120)
		fmt.Printf("%v\n", err)
		val, found, err := GetValue("sec-key")
		fmt.Printf("%v %v %v\n", val, found, err)
	} else {
		fmt.Println("Error in InitRedis: ", err.Error())
	}

}
