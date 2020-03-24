package main

import (
	"fmt"
)

func test() {
	defer fmt.Println("----------1-------")
	//runtime.Goexit()   //直接结束go程
	//return 返回调用者
	fmt.Println("--------2--------")
}

func main() {
	go func() {
		defer fmt.Println("----------3-------")
		go test()
		fmt.Println("--------4--------")

	}()
	for {

	}
}
