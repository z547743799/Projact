package main

import (
	"fmt"
	"os"
)

func main() {
	//就是输出打印
	fmt.Fprintf(os.Stdout, "[%s]\n", "Yinzhengjie") //fmt.Fprintf方法要求第一个参数必须有可写的方法，第二个参数用于定义格式化输出，第三个参数是具体的字符串。
	//buf := bufio.NewWriter(os.Stdout)               //我们创建一个写入器名字叫做buf，注意，bufio.NewWriter方法需要传入的对象必须有一个可写的方法。
	//fmt.Fprintf(buf, "【%s】\n", "Golang")
	//buf.Flush() //在缓冲写入的最后千万不要忘了使用 Flush() ，否则最后的输出不会被写入。
}
