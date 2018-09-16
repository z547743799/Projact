package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	fmt.Printf("Get mykey: %v \n", username)
	s := time.Now()
	s5, _ := strconv.Atoi(username)

	fmt.Println(s5)
	if s.Unix() >= int64(s5) {
		fmt.Println(true)
	}
}
