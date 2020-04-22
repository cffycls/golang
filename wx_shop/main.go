package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"reflect"
	"time"
)

const (
	REDIS_KEY = "mobile:p40:n3000:1"
)
func main() {
	client := redis.NewClient(&redis.Options{
		Addr:               "172.1.13.11:6379",
		Password:           "",
		DB:                 0,
		DialTimeout:		time.Second * 2,
	})
	defer client.Close()
	length,_ := client.LLen(REDIS_KEY).Result()
	fmt.Println(length, reflect.TypeOf(length).String(), length > 0)
	if length != 10{
		client.LTrim(REDIS_KEY, 1,0) //清空
		vals,_ := client.LRange(REDIS_KEY,0,-1).Result()
		println(reflect.TypeOf(vals).String(), "\n??----")
		fmt.Println(vals, len(vals))
		for i:=0; i<len(vals); i++ {
			fmt.Print(vals[i], "  ")
		}
		fmt.Println("----??")
		client.LPush(REDIS_KEY, 1,2,3,4,5,6,7,8,9,10)
	}

	vals,_ := client.LRange(REDIS_KEY,0,-1).Result()
	println(reflect.TypeOf(vals).String())
	fmt.Println(vals, len(vals))
	for i:=0; i<len(vals); i++ {
		fmt.Print(vals[i], "  ")
	}
	fmt.Println()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}