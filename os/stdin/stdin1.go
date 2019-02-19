package stdin

import (
	"bufio"
	"fmt"
	"os"
)

func Stdin1() {
	//读取输入
	reader := bufio.NewReader(os.Stdin)

	//将输入转化为string
	result, err := reader.ReadString('\n')
	if err != nil {

		fmt.Println("read error:", err)
	}

	fmt.Println("result:", result)
}
