package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 1000)
	}
}
func person1() {
	printer("hello")
	ch <- 789
}

func person2() {
	<-ch
	printer("world")

}

func main() {
	go person1()
	go person2()
	for {

	}
}
