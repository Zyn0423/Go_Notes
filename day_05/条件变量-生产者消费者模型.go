package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond // 定义全局变量

func producer09(out chan<- int, idx int) {
	for {
		cond.L.Lock()

		for len(out) == 3 {
			cond.Wait() //1.阻塞2.解锁3.释放
		}
		num := rand.Intn(800)
		out <- num
		fmt.Printf("----%dth 写 go程，写出：%d\n", idx, num)
		cond.L.Unlock() //解锁
		cond.Signal()   //唤醒
		time.Sleep(time.Millisecond * 1000)
	}

}

func consumer09(in <-chan int, idx int) {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait() //1.阻塞2.解锁3.释放
		}
		num := <-in
		fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 1000)
	}

}

func main() {
	//consumer09 消费者
	//producer09 生产者

	ch := make(chan int, 3)
	quit := make(chan bool)
	rand.Seed(time.Now().UnixNano())
	cond.L = new(sync.Mutex) //ToDo 指定条件使用的锁
	for i := 0; i < 5; i++ {
		go producer09(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go consumer09(ch, i+1)
	}
	<-quit
}
