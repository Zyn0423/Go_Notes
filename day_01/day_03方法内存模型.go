package main

import "fmt"

type student struct {
	Name  string
	Age   int
	Adder string
}

//TODO 方法调用中 方法类型是指针类型
func (s *student) work() {
	fmt.Println(s.Age)
	s.Age = 80
	fmt.Println(s)
}
func main() {
	//地址传递
	s := &student{"亚索", 20, "德州"}
	s.work()
	fmt.Println(*s)

}
