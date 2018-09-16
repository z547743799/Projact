package main

import (
	"bytes"
	"fmt"

	"github.com/cuixin/go/codec"
	"github.com/garyburd/redigo/redis"
)

type test struct {
	strattime int64
	endtime   int64
}

var h = new(codec.JsonHandle)
var map1 = make(map[int]*test)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	s := new(test)
	s.strattime = 1
	s.endtime = 123
	a := test{strattime: 1, endtime: 23}
	fmt.Println(s, a)
	map1[0] = s
	map1[1] = &a

	data, _ := Encode(map1)
	_, err = c.Do("SET", "key", data)

	for i := 0; i < 2; i++ {
		fmt.Println(map1[i])
	}
}
func Encode(obj interface{}) ([]byte, error) {
	w := make([]byte, 0, 2048)
	writer := bytes.NewBuffer(w)
	enc := codec.NewEncoder(writer, h)
	err := enc.Encode(obj)
	return writer.Bytes(), err
}
