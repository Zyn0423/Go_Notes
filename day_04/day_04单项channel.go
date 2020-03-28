package main

import "fmt"

func main_01() {

	ch := make(chan int) //TODO 双向chan

	var write_ch chan<- int = ch

	write_ch <- 20 // TODO 单向chan
	fmt.Println(write_ch)
	//		TODO  fatal error: all goroutines are asleep - deadlock!  使用chan 必须两端同时进行
	var read_ch <-chan int = ch //TODO 单向读chan

	num := <-read_ch
	fmt.Println(num)

}

func write_ch(ch chan<- int) {
	ch <- 77
}

func read_ch(ch <-chan int) {
	num := <-ch
	fmt.Println("读端", num)
}

func main() {
	ch := make(chan int) //定义一个双向chan
	go func() {
		write_ch(ch) // TODO 传引用
	}()
	read_ch(ch)

}
