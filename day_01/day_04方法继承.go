package main

import "fmt"

type person01 struct {
	Name string
	Age  int
	id   int
}

type student02 struct {
	person01 // 匿名字段
	class    string
}

func (a *person01) info() {
	fmt.Println(*a)
}

func main() {
	a := &person01{"亚索", 20, 11}
	a.info()
	//TODO 子类可以继承父类  ,但是父类不能继承子类不管属性和方法都不行
	//TODO 写方法的时候建议使用指针类型
	b := &student02{person01{"德玛", 20, 50}, "纽约"}
	b.info()
}
