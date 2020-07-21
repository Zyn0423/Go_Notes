package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//每首先21万块减半
	//最初奖励50BTC
	//用一个循环来判断、增加
	total := 0.0
	blockInterval := 21.0
	currentReward := 50.0
	for currentReward > 0 {
		//每一个区块的总量
		amount1 := blockInterval * currentReward
		currentReward *= 0.5
		total += amount1
	}
	fmt.Println("比特币总量", total, " 万")
}
