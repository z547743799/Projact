package main

import "fmt"

type Person5 struct {
	name string
	age  int
}

//在结构体中属于匿名结构体的字段称为提升字段
type Student struct {
	Person5
	school string
}

func (s *Student) Animal() {
	fmt.Println(s.name, "---")
}

func main() {

	p1 := Person5{"王二狗", 30}
	fmt.Println(p1.name, p1.age)

	s1 := Student{}
	s1.name = "李小花" // 提升字段
	s1.age = 30     // 提升字段
	s1.school = "清华小班"
	s1.Animal()
	fmt.Println(s1.name, s1.age, s1.school)
	fmt.Println(s1.age, s1.name)
	ss := &p1
	fmt.Println(ss, "---")
	sw := ss
	fmt.Println(sw, "---")

	s2 := Student{Person5{"如梦", 20}, "清华大班"}
	fmt.Println(s2)
	s3 := Student{Person5: Person5{"如歌", 23}, school: "清华中班"}
	fmt.Println(s3)

}
