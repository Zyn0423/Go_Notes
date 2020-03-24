package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1)
	fmt.Println("n =", n)

}
