package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	user := "23423"
	md := md5.New()
	md.Write([]byte(user))
	use := md.Sum(nil)
	fmt.Printf("%x", use)
}
