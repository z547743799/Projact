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

	t1 := time.Now()
	fmt.Println(t1)
	m1, _ := time.ParseDuration("+10s")
	m2 := t1.Add(m1)
	fmt.Println(m2)
	fmt.Println(m2.Unix() - t1.Unix())
	fmt.Println(t1.Unix())
	fmt.Println(m2.Unix())
	sd := m2.Unix()
	s5 := strconv.Itoa(int(sd))
	_, err = c.Do("SET", "mykey", s5)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

}
