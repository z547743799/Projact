package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/ugorji/go/codec"
)

type Ss struct {
	strattime int64
	endtime   int64
}

var h = new(codec.JsonHandle)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	username, err := redis.String(c.Do("GET", "key"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	fmt.Printf("Get mykey: %v \n", username)

	ss := new(Ss)
	err = Decoder([]byte(username), ss)
	if err != nil {
		fmt.Println(err)
		return
	}

	s := time.Now()

	fmt.Println(ss.endtime, "===")
	if s.Unix() >= ss.endtime {
		fmt.Println(true)
	}

}
func Decoder(data []byte, v interface{}) error {
	reader := bytes.NewBuffer(data)
	decoder := codec.NewDecoder(reader, h)
	err := decoder.Decode(v)
	return err
}
