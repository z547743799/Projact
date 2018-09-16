package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//全局变量
var ticket2 = 10 // 100张票
var wg3 sync.WaitGroup

var matex sync.Mutex // 创建锁头

func main() {
	/*
		4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	*/
	wg3.Add(4)
	go saleticket2s("售票口1") // g1,100
	go saleticket2s("售票口2") // g2,100
	go saleticket2s("售票口3") //g3,100
	go saleticket2s("售票口4") //g4,100
	wg3.Wait()              // main要等待。。。
}

func saleticket2s(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg3.Done()
	//for i:=1;i<=100;i++{
	//	fmt.Println(name,"售出：",i)
	//}

	for { //1, g1 ,g2,g3,g4
		//上锁， g2
		matex.Lock()
		if ticket2 > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket2) // 1 , 0, -1 , -2
			ticket2--                         //0 , -1 ,-2 , -3
		} else {
			matex.Unlock() //条件不满足，也要解锁，结束循环
			fmt.Println("售罄，没有票了。。")
			break
		}
		matex.Unlock() //解锁
	}
}
