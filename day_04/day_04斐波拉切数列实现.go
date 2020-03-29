package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			runtime.Goexit()
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	x, y := 1, 1           //   x=1  y=1
	go fibonacci(ch, quit) //TODO 读数据

	for i := 0; i < 10; i++ { //TODO 写数据
		ch <- x
		x, y = y, x+y
	}
	quit <- true // todo 结束go程

}
