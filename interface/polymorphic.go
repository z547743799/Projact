package main

import (
	"fmt"
)

type Run interface {
	run()
}
type Animal struct {
	school string
	sed    interface{}
}

type Runner struct {
	Animal
	name string
}

func (r *Runner) run() {
	fmt.Println(r.school, r.name, "，开始跑步。。。慢跑")
}

func main() {
	var r1 Run // 对象的声明
	ssw := &Runner{}
	ssw.school = "232"
	ssw.name = "sdf"
	if v, ok := ssw.sed.(string); ok {
		fmt.Println(v)
	}
	r1 = ssw
	r1.run()
	ss := "sdf"
	if s, ok := ss.(string); ok {
		fmt.Println(s)
	}

}
