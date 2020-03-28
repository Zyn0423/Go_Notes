package main

import "fmt"

type OrderInfo struct {
	id int
}

func producer2(out chan<- OrderInfo) {
	for i := 0; i < 10; i++ {
		fmt.Println("订单 ", i+1)
		order := OrderInfo{id: i + 1}
		out <- order
	}
	close(out)
}

func consumer2(in <-chan OrderInfo) {
	for num := range in {
		fmt.Println("读数据", num.id)
	}
}
func main() {
	ch := make(chan OrderInfo, 5) //异步    、
	//ch :=make(chan OrderInfo)  //同步
	go producer2(ch)
	consumer2(ch)
}
