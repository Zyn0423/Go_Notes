package main

import "fmt"

func test001(n int) {
	defer func() { // TODO 使用defer 延迟调用 匿名函数
		if err1 := recover(); err1 != nil {
			fmt.Println(err1)
		}
	}()
	var p [3]int
	p[n] = 111
	fmt.Println(n)
}
func test002() {

	fmt.Println("test002")
}
func test003() {
	fmt.Println("test003")
}
func main() {
	test002()
	test001(1)
	test003()
}
