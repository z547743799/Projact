package main

import (
	"Project/xiaoyao/codec/code"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Message2 struct {
	Starttimes int64
	Endtimes   int64
}

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

	ss := Message2{}

	err = code.Decoder([]byte(username), &ss)

	if err != nil {
		fmt.Println(err)
		return
	}

	s := time.Now()

	fmt.Println(ss.Endtimes, "===")
	if s.Unix() >= ss.Endtimes {
		fmt.Println(true)
	}

}
