package main

import (
	"Project/xiaoyao/codec/code"
	"fmt"
	"reflect"
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

	t1 := time.Now()
	fmt.Println(t1)
	m1, _ := time.ParseDuration("+20s")
	m2 := t1.Add(m1)
	fmt.Println(m2)
	fmt.Println(m2.Unix() - t1.Unix())
	fmt.Println(t1.Unix())
	fmt.Println(m2.Unix())
	//ss := Ss{1221, 12}
	s1 := Message2{t1.Unix(), m2.Unix()}
	//	ss.strattime = t1.Unix()
	//	ss.endtime = m2.Unix()
	//  s5 := strconv.Itoa(int(sd))
	fmt.Println(m2.Unix())
	//fmt.Println(ss)

	data, _ := code.Encode(s1)
	if err != nil {
		fmt.Println("redis set", err)
	}
	fmt.Println(data, "00101001")
	a1 := reflect.TypeOf(data)
	fmt.Println(a1)

	_, err = c.Do("SET", "mykey", data)

	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}
