package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i

		}
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		num, ok := <-ch
		if ok == true {
			fmt.Println("读数据", num)
		} else {
			break

		}
	}
}
