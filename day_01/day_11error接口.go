package main

import "fmt"

func test(a, b int) (n int, err error) {
	if b == 0 {
		fmt.Println(err)
	} else {
		n = a / b
	}
	return
}
func main() {
	n, err := test(10, 0)
	if err != nil {
		fmt.Println("....")
	}
	fmt.Println(n)
}
