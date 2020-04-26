package main

import "fmt"

type preson5 struct {
	Name  string
	Age   int
	Adder string
}

type bigpreson struct {
	preson5
}

//TODO 先定义接口一般er 结尾  根据接口实现功能
type Humber interface {
	//TODO 声明方法
	work()
}

func (s *preson5) work() {
	fmt.Printf("我叫%s ,今年%d ,我住%s ", s.Name, s.Age, s.Adder)
}
func (s *bigpreson) work() {
	fmt.Printf("我叫%s ,今年%d ,我住%s 哈哈", s.Name, s.Age, s.Adder)
}
func main() {
	//TODO 接口是一种数据类型 ，可以接受满足对象的信息
	// 接口是虚的方法是实的
	// 接口定义的规则，在方法中有定义的实现

	var h Humber
	s := preson5{"德玛西亚", 20, "德玛大陆"} //
	//s:=bigpreson{preson5{"德玛西亚",20,"德玛大陆"}}
	//讲对象赋值给接口类型变量
	h = &s
	h.work()
	//s.work()

}
