package test

import (
	"fmt"
	"io/ioutil"
)

//递归打印目录
func Abelive(a string) {

	fileInfos, _ := ioutil.ReadDir(a)
	for _, fi := range fileInfos {
		filename := a + "/" + fi.Name()
		fmt.Println(filename)
		if fi.IsDir() {
			//继续遍历fi这个目录
			Abelive(filename)
		}
	}
}

//rune 转byte
func Belive(ss rune) []byte {
	robot := new(Robot)
	robot.Work()
	sd := make([]byte, 0)
	e := byte(ss)
	se := append(sd, e)
	return se
}

type Robot struct {
	action int
}

func (s *Robot) Work() {
	s.action = 1
	fmt.Println("work", s.action)
}

func (s *Robot) Speak() {
	fmt.Println("speak")
}

func (s *Robot) Look() {
	fmt.Println("look")

}
