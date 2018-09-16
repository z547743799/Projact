package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Println("//获取时间")
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println()
	fmt.Println("//减两个小时")
	s1, _ := time.ParseDuration("+2h")
	s2 := t1.Add(s1)
	fmt.Println(s2)
	fmt.Println("//两个时间戳")
	fmt.Println(t1.UnixNano())
	fmt.Println(s2.UnixNano())
	fmt.Println(t1.Unix() - s2.Unix())

	fmt.Println("//获取10秒后的时间判断当前时间是否大于时间戳")
	m1, _ := time.ParseDuration("+10s")
	m2 := t1.Add(m1)
	fmt.Println(m2)
	fmt.Println(m2.Unix() - t1.Unix())
	fmt.Println(t1.Unix())
	fmt.Println(m2.Unix())

	if t1.Unix() > 1536566753 {
		fmt.Println(true)

	}

}
