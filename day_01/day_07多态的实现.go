package main

import "fmt"

type opterate struct {
	num1 int
	num2 int
}
type Add struct {
	opterate
}
type sub struct {
	opterate
}

func (s *Add) work() int {
	return s.num1 + s.num2
}
func (s *sub) work() int {
	return s.num1 - s.num2
}

type Opter interface {
	work() int
}

func calculate(s Opter) {

	fmt.Println(s.work())
	//fmt.Println(s)
}
func main() {
	//var o Opter
	s := Add{opterate{1, 2}}
	calculate(&s)
	b := sub{opterate{3, 2}}
	calculate(&b)
	//o=&s
	//fmt.Println(o.work())
	//b :=sub{3,2}
	//o =&b
	//fmt.Println(	o.work())
}
