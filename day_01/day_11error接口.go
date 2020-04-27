package main

import "fmt"

func test(a, b int) (n int, err error) {
	err = nil
	if b == 0 {

	} else {
		n = a / b
	}
	return
}
func main() {
	n, err := test(10, 0)
	if err != nil {
		fmt.Println("....", err)
	}
	fmt.Println(n)
}
