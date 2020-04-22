package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

func main()  {
	client := redis.NewClient(&redis.Options{
		Addr:               "172.1.13.11:6379",
		Password:           "",
		DB:                 0,
		DialTimeout:		time.Second * 2,
	})
	defer client.Close()

	key := "mobile:p40:n3000:1"

	length,_ := client.LLen(key).Result()
	if length > 0 {
		for i:=0; i<int(length); i++ {
			v,_ := client.LPop(key).Result()
			fmt.Println(v)
		}
	}

}
