package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex // 锁只有一把， 2 个属性 r w
func read_go(in <-chan int, idx int) {
	for {
		rwMutex.RLock() // 以读模式加锁
		num := <-in
		fmt.Printf("----%dth 读go程 读出: %d\n", idx, num)
		rwMutex.RUnlock() // 以读模式解锁
	}
}
func write_go(out chan<- int, idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex.Lock() // 以写模式加锁
		out <- num
		fmt.Printf("----%dth 写go程 写出: %d\n", idx, num)
		time.Sleep(time.Millisecond * 1000)
		rwMutex.Unlock()
	}
}
func main() {

	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go read_go(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go write_go(ch, i+1)
	}

	for {

	}
}
