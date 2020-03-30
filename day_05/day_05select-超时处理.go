package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() { //TODO 子go 程获取数据
		for {
			select {
			case num := <-ch:
				fmt.Println("读", num)
			case <-time.After(time.Second * 3): // TODO 时间定时器
				quit <- true
				//goto lable //TODO 标签跳转
				return

			}
		}
		//lable:
	}()

	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second * 2)

	}
	<-quit //  结束

}
