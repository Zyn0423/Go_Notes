package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)                   //TODO 创建Channel  为了配合go程的结束
	fmt.Println("现在的时间", time.Now())        //获取现在的时间
	now_time := time.NewTicker(time.Second) //	C <-chan Time // The channel on which the ticks are delivered.i
	go func() {
		i := 0 // 创建一个循环计数器
		for {
			now_ing := <-now_time.C //实例化该方法
			fmt.Println(now_ing)
			i++ //变量累加
			if i == 3 {
				ch <- true
				break //return runtime.exit()

			}

		}
	}()
	<-ch

}
