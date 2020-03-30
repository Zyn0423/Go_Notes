package main

import (
	"fmt"
	"math/rand"
	"time"
)

func write_go2(in chan<- int, idx int) {
	for {
		in <- rand.Intn(1000)
		fmt.Printf("----%dth 写 go程，写出：%d\n", idx, in)
		time.Sleep(time.Millisecond * 300)
	}

}
func read_go2(out <-chan int, idx int) {
	for {
		num := <-out
		fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		go write_go2(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go read_go2(ch, i+1)
	}

	for {

	}
}
