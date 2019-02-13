package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	result, err := reader.ReadString('\n')
	if err != nil {

		fmt.Println("read error:", err)
	}

	fmt.Println("result:", result)

}
