package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		获取日期，时间
			ctrl+f,ctrl+r
	*/
	t1 := time.Now()
	fmt.Println(t1)
	//获取日期
	year, month, day := t1.Date()
	fmt.Println(year, "年", month, "月", day, "日")
	//获取时间
	hour, minute, second := t1.Clock()
	fmt.Println(hour, "时", minute, "分", second, "秒")
	//获取年，月，日，时，分，秒，星期，纳秒，。。
	fmt.Println("年：", t1.Year())        //年
	fmt.Println("月：", t1.Month())       //月
	fmt.Println("日：", t1.Day())         //日
	fmt.Println("时：", t1.Hour())        //时
	fmt.Println("分：", t1.Minute())      //分
	fmt.Println("秒：", t1.Second())      //秒
	fmt.Println("纳秒：", t1.Nanosecond()) //纳秒
	fmt.Println("星期：", t1.Weekday())    //星期
	fmt.Println(t1.ISOWeek())           //返回年份，第几周
	fmt.Println(t1.YearDay())
}
