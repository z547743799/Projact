package main

import (
	"fmt"
	"time"
)

func main() {
	//1.获取时间对象
	t1 := time.Now() // 当前的时间
	fmt.Println(t1)  //2018-04-12 16:36:11.729 +0800 CST m=+0.009000001
	fmt.Println(t1.Hour() - 2)
	fmt.Println(t1)

	fmt.Println(t1.Unix())
	//创建指定的时间
	t2 := time.Date(1998, 4, 12, 16, 34, 0, 0, time.Local)
	fmt.Println(t2)

	/*
			 time-->string
		 	 	t1.Format("格式")-->string
				 string-->time
			 	 	time.Parse("格式","数据字符串")-->time
	*/
	s1 := t1.Format(time.ANSIC)                 //Mon Jan _2 15:04:05 2006
	s2 := t1.Format("Mon Jan _2 15:04:05 2006") //6-1-2-3-4-5
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("--------")

	s3 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s3)
	s4 := t1.Format("2006-1-2 15:04:05")
	fmt.Println(s4)

	src1 := "1998年4月12日 17:33:24"
	t3, err := time.Parse("2006年1月2日 15:04:05", src1)
	fmt.Println(t3, err)

	/*
			  3.timestamp,时间戳：
		  	  		日期，距离1970年1月1日，0点0时，0分0秒，时间的差值
	*/
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	//1970年1月1日，1点0分0秒
	//1小时，60分钟，3600秒 3600 000 000 000
	timeStamp1 := t4.Unix()     //秒 差值 3600
	timeStamp2 := t4.UnixNano() //纳秒
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)

	fmt.Println(t1.Unix())     //当前时间，距离指定日期的秒数
	fmt.Println(t1.UnixNano()) //当前时间，距离指定日期的纳秒数
}
