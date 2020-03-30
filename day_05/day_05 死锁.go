package main

import "fmt"

func main_1() {
	ch := make(chan int) //all goroutines are asleep - deadlock!
	num := <-ch

	ch <- 789
	fmt.Println(num)

}

func main_2() {
	ch := make(chan int)
	go func() {
		ch <- 789
	}()

	num := <-ch
	fmt.Print(num)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num

			}
		}

	}()
	for { //TODO fatal error: all goroutines are asleep - deadlock!
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}

}
