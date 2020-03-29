package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)    //TODO  用来数据通信的channel
	quit := make(chan bool) //TODO 用来判断是否退出channel

	go func() { //TODO 写数据
		for i := 0; i < 5; i++ {
			ch <- i
		}
		quit <- true     //TODO 通知主GO程退出
		runtime.Goexit() //TODO 退出子go程
	}()

	for {
		select { // TODO 主GO程读数据
		case num := <-ch:
			fmt.Println("数据流", num)
		case <-quit:
			return //TODO 返回主go程     如果要用break 只是跳出select语句，for循环还在运行
		}

	}
}
