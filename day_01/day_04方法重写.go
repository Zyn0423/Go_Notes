package main

import (
	"fmt"
)

type person03 struct {
	Name string
	Age  string
}

type student03 struct {
	person03
	Adder string
}

func (s *person03) Printinfo() {
	fmt.Println(*s)
}
func (s *student03) Printinfo() {
	fmt.Println(*s)
}
func main() {
	s := student03{person03{"武器大师", "20"}, "天空之城"}

	s.Printinfo()          // TODO 子类对象方法，采用就近原则，使用子类方法
	s.person03.Printinfo() //TODO 父类对象方法

}
