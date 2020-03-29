package main

import (
	"fmt"
	"time"
)

func main_1() {
	fmt.Println(time.Now()) //获取当前时间

	//1. sleep
	//time.Sleep(time.Second*2)
	//2.Time.C
	//new_time:=time.NewTimer(time.Second*2)  //	C <-chan Time
	//now_time1:=<-new_time.C   //返回单项chan 写，要接收读操作
	//fmt.Println("现在时间",now_time1)
	//3.time.After
	//now_time:=time.After(time.Second*2)
	//fmt.Println(<-now_time)
}

func main() {
	new_time := time.NewTimer(time.Second * 3) //	C <-chan Time
	new_time.Reset(time.Second * 1)            //TODO 重置
	go func() {

		<-new_time.C //
		fmt.Println("读")
	}()
	new_time.Stop() //TODO  停止，  会阻塞
	for {

	}

}
