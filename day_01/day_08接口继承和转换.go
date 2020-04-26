package main

import "fmt"

type person06 struct {
	Name string
	Age  int
}
type H1umber interface { // TODO 子集
	sayhi()
}

func (s *person06) sayhi() {
	fmt.Println(s.Name, s.Age)
}

func (s *person06) singder(count string) {
	fmt.Printf("%s 次统计", count)
}

type H011umber interface { //TODO  超集
	H1umber // 匿名字段
	singder(count string)
}

//func singder(count string)  {
//	fmt.Printf("%s 次统计")
//}

func main() {
	var h H1umber
	var h2 H011umber
	s := person06{"蛮王", 20}
	// TODO 子集不包含超集，不能讲子集赋值给超集 但是超集包含子集能把超集赋值给子集
	h2 = &s
	//h.sayhi()
	h = h2
	//h.sayhi()

	h.sayhi()
	//h.singder("20")

}
