package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子GO程,i", i)
			ch <- i
		}
	}()
	//TODO 打印屏幕属于IO操作，队列等待，在等待时切换到子go程
	//TODO 如果加等待时间，到主Go程时没有涉及到CPU，顺利的获得访问屏幕的权力  （程序调度CPU来控制 ，它调用不调用看能否争夺到时间轮片这种争夺是随机 ）
	time.Sleep(time.Second * 1)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主go程", num)
	}

}
