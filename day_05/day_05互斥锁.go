package main

import (
	"fmt"
	"sync"
	"time"
)

//var ch = make(chan int)
//
//func printer(str string) {
//	for _, ch := range str {
//		fmt.Printf("%c", ch)
//		time.Sleep(time.Millisecond * 1000)
//	}
//}
//func person1() {
//	printer("hello")
//	ch <- 789
//}
//
//func person2() {
//	<-ch
//	printer("world")
//
//}
//
//func main() {
//	go person1()
//	go person2()
//	for {
//		;
//	}
//}

//TODO  使用锁完成同步--互斥锁
var mutex sync.Mutex // TODO创建一个互斥量,新建的互斥锁状态为0未加锁，锁只有一把

func printer(str string) {
	mutex.Lock() // TODO 访问共享数据之前，加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 100)
	}
	mutex.Unlock() //TODO 共享数据访问结束，解锁
}
func person1() {
	printer("hello") // 先
}

func person2() {
	printer("world") //后

}

func main() {
	go person1()
	go person2()
	for {

	}
}
