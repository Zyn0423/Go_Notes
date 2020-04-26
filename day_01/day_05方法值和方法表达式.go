package main

import "fmt"

// TODO 方法值需要绑定实例
//TODO   方法表达式显示传参

type person04 struct {
	Name  string
	Age   int
	Adder string
}

type skills struct {
	person04
	context_ string
}

func (s *skills) Printinfo() {
	fmt.Println(*s)
}
func main() {
	//TODO 隐式传递        方法值    隐藏的是接受者绑定是对象
	s := &skills{person04{"剑圣", 20, "安德鲁"}, "疾风"}
	//s1 :=s.Printinfo
	//s1()

	//显示传参     方法表达式
	//TODO 对象相同函数类型相同不能赋值
	//TODO 函数类型相同可以赋值
	//s1:=(*skills).Printinfo
	//s1(s)

	(*skills).Printinfo(s)

	//fmt.Println(s1)

}
