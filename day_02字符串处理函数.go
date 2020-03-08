package main

import (
	"fmt"
	"strings"
)

func testSplit() {
	str := "Now that we have a feel for what the data elements look like, how do we manage them with our application?"

	ret := strings.Split(str, " ") //TODO 字符串按指定分割返回切片类型
	fmt.Println(ret)
	for _, s := range ret { //循环遍历的当时
		fmt.Println(s)
	}
	ret2 := strings.Fields(str) // TODO 字符串按空格拆分
	fmt.Println(ret2)

}

func hasfuffix_preffix() { // TODO 判断字符串结束标记和起始标记

	result := strings.HasSuffix("Happybirthday.mp3", "mp3") //结束
	fmt.Println(result)
	result2 := strings.HasPrefix("Happybirthday.mp3", "Happybirthday") //起始
	fmt.Println(result2)

}

func main() {
	testSplit()
	hasfuffix_preffix()
}
