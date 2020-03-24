package main

import (
	"fmt"
	"time"
)

var channel1 = make(chan int)

func Printer(s string) {
	for _, v := range s {

		fmt.Printf("%c", v) //TODO 输出单个字符。
		time.Sleep(1000 * time.Millisecond)
	}
}

func Person1() {
	Printer("hello")
	channel1 <- 8
}
func Person2() {
	<-channel1
	Printer("world")
}
func main() {
	go Person1()
	go Person2()
	for {

	}

}
