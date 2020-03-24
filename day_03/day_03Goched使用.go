package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for {
			fmt.Println("-------Goroutine------")
			//time.Sleep(time.Second)
		}
	}()
	for {
		runtime.Gosched() //TODO 出让当前CPU时间
		fmt.Println("-----Pro--------")
	}
}
