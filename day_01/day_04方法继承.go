package main

import "fmt"

type person01 struct {
	Name string
	Age  int
	id   int
}

type student02 struct {
	person01
	class string
}

func (a *person01) info() {
	fmt.Println(*a)
}

func main() {
	a := &person01{"亚索", 20, 11}
	a.info()

	b := &student02{person01{"德玛", 20, 50}, "纽约"}
	b.info()
}
