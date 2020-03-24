package main

import (
	"fmt"
	"time"
)

func sing() {

	for i := 0; i < 5; i++ {
		fmt.Println("-------正在唱歌----------")
		time.Sleep(1000 * time.Millisecond)
	}

}

func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("-------正在跳舞----------")
		time.Sleep(1000 * time.Millisecond)
	}
}
func main() {
	go sing()
	go dance() //TODO 子任务
	//TODO 父进程
	for {

	}
}
