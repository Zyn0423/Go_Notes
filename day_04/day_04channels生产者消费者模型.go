package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("写入数据", i*i)
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("读到数据", num)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch) //TODO 子go程 生产者

	consumer(ch) // TODO 消费者

}
