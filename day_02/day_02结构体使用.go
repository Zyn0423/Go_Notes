package main

import (
	"fmt"
	"unsafe"
)

// TODO 结构体：是一种数据类型

type Person struct { //类型定义  不能给成员里面赋值的（地位等价于 int byte boll string...）
	name string
	age  int
	sex  byte
}

type user string //自定义数据类型

//TODO 普通变量定义和初始化
//1. 顺序初始化
var man = Person{"张三", 20, 'm'} //全员必须初始化

//2.指定成员初始化：

//3.结构体函数
func test1(temp Person) {
	fmt.Println(unsafe.Sizeof(temp)) //TODO 查看结构体大小 返回的数字
	temp.name = "王麻子"
	fmt.Println(temp)
}
func main_01() {
	man2 := Person{name: "李四"}
	fmt.Println(man)
	fmt.Println(man2)
	fmt.Printf("%q\n", man.name) //索引成员变量
	var man3 Person
	man3.name = "王五"
	man3.age = 20
	man3.sex = 'f'
	fmt.Printf("%q\n", man3)
	// TODO 机构体比较  只能用 == ！=     不能用<>
	var p1 = Person{"张三", 20, 'm'}
	var p2 = Person{"张三", 20, 'm'}
	var p3 = Person{"张三", 201, 'm'}
	fmt.Println("p1 == p2", p1 == p2)
	fmt.Println("p1 ==p3", p1 == p3)

	//	TODO 相同类型结构体赋值（变量类型，个数，顺序一致） 不相同的不可以
	// TODO 结构体变量的地址 == 结构体首个元素的地址
	var tmp Person
	tmp = p3
	fmt.Println(tmp)

	//TODO 结构体传参：将结构体变量的值拷贝一份，传递。--（内存消耗大）
	var tmep Person
	test1(tmep)
}

//TODO 指针结构体定义和使用
type Person1 struct {
	name string
	age  int
	sex  byte
}

func test3(data *Person1) {
	data.name = "高"
	data.age = 50
	data.sex = 'm'

}
func main3() {
	//	TODO 1.第一种方法
	var p1 = &Person1{name: "王", age: 20, sex: 'f'}

	fmt.Println(*p1)
	fmt.Printf("p1 = %T\n", p1)
	//TODO 第二种方法

	p2 := new(Person1)
	p2.name = "李"
	p2.sex = 'm'
	p2.age = 50
	fmt.Println("new", *p2)

	//TODO 结构体函数传参

	fmt.Printf("%T\n", p2)
	test3(p2)
	fmt.Println(p2)

}

//TODO 结构体指针做函数的返回值

type Person4 struct {
	name string
	age  int
	sex  byte
	ids  []string
}

func init_print(data *Person4) {
	data.name = "1231"
	data.sex = 'n'
	data.age = 20
	data.ids = append(data.ids, "123", "456")
	//fmt.Println(data)

}
func init_print2() *Person4 {
	//var data Person4
	data := new(Person4)
	data.name = "666"
	data.sex = 'n'
	data.age = 15
	data.ids = append(data.ids, "123", "456")
	//TODO 返回局部变量的值， 而&data是返回局部变量的地址值的
	return data
}

func main() {
	var p1 Person4 //TODO 初始值函数内部传普通结构体地址的方法
	init_print(&p1)
	//fmt.Println(p1)
	//p2 := init_print2() //TODO 初始值带返回值的方法 注意：不能返回局部变量的地址值，原因： 局部变量保存在栈针上，函数调用结束后，会立即释放，但是不会把值清除掉，但不会受系统保护，随时可能分配到其他程序
	//fmt.Println(*p2)
	//fmt.Printf("%T\n", p2)

}
