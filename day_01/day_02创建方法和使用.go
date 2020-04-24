package main

import "fmt"

//func (对象)方法(参数列表)返回值列表{代码体}
type cat struct {
	Name string
	Age  int
}
type dog struct {
	Name string
	Age  int
}

//todo 结构体当中对象来传递
func (d dog) show() {
	fmt.Println("喵喵喵")
}

//todo 方法名一样接受者不一样
func (d cat) show() {
	fmt.Println("汪汪汪")
}

func main() {
	var d dog
	d.Name = "m"
	d.Age = 10
	d.show()
	fmt.Print(d)
	//TODO 对象方法名和函数名可以重复，但是相同的对象方法不能重名
	var m cat
	m.Name = "x"
	m.Age = 12
	m.show()

}
