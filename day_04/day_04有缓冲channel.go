package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Println("ch", len(ch), cap(ch))
	go func() {
		for i := 0; i < 5; i++ {
			//fmt.Println("子go程i=", i)
			ch <- i
			fmt.Println("子go程i=", i)
		}
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("num", num)
	}
}
