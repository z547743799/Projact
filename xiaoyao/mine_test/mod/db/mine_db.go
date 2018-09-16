package db

import (
	"Project/xiaoyao/codec/code"
	"Project/xiaoyao/mine_test/data"
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

func (this *DbMod) GetMinerConfig() map[int64]*data.TMineConfig {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}
	defer c.Close()

	mr := make(map[int]*data.TMineCaveConfig)

	reply, err := redis.StringMap(c.Do("hgetall", "pool"))
	for k, v := range reply {
		key, err := strconv.Atoi(k)
		if err != nil {
			fmt.Println("0101010101010", err)
		}
		var r data.TMineCaveConfig
		err = code.Decoder([]byte(v), &r)
		if err != nil {
			fmt.Println("----------", err)
		}
		mr[key] = &r
	}
	return nil

}
