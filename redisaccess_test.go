package redisaccess

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	InitRedis("localhost", "")
	err := SetValue("sec-key", "My Data", time.Second*120)
	fmt.Printf("%v\n", err)
	val, found, err := GetValue("sec-key")
	fmt.Printf("%v %v %v\n", val, found, err)

}
