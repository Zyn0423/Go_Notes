package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "123.433 a7c.4 mfc cat aMc azc cba"

	regexp1, err := regexp.Compile(`\d+\.\d+`)
	if err != nil {
		fmt.Println("regexp.Compile", err)
	}
	result := regexp1.FindAllStringSubmatch(str, -1)
	fmt.Println(result)

}
