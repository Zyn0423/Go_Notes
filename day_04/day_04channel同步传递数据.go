package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	fmt.Println("i=", ch, "len(chan)", len(ch), "cap(chan)", cap(ch))
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i=", i, "len(chan)", len(ch), "cap(chan)", cap(ch))

		}
		ch <- "打印完成"
	}()

	rest := <-ch
	fmt.Println(rest)
	time.Sleep(1000 * time.Second)
}
