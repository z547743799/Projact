package main

import (
	"fmt"
	"os"
	"os/exec"
)

//1+2
func main() {
	//写入
	fileName := "son.go"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s := "package main"
	dstFile.WriteString(s + "\n \nimport \"fmt\"\n \nfunc main(){\n\n fmt.Println(\"vim-go\")\n\n}")
	//执行
	command := `go run son.go`
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
