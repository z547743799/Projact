package main

import "fmt"
import "os"

func main() {
	fileName := "son.go"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s := "package main"
	dstFile.WriteString(s + "\n\nimport \"fmt\"\n \nfunc main(){\n\n fmt.Println(\"vim-go\")\n\n}")
}
