package main

import "fmt"

func retsult() {
	fmt.Println("func go run ...")
}
func main() {
	var p []interface{} // TODO 接口可以接受任意类型

	p = append(p, "123", 123, retsult)
	//g:=p[2]
	s := p[2]
	fmt.Println(s)
	fmt.Println(&s)
	fmt.Printf("%T", s)

}
