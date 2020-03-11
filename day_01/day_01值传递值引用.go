package main

import (
	"fmt"
	"strings"
)

func swap(a, b int) {
	a, b = b, a
	fmt.Println("a=", a, "b=", b)
}
func swap_1(a, b *int) {
	*a, *b = *b, *a //*解引用
	fmt.Println("a=", a, "b=", b)
}

func slice_1() {
	//练习1：给定一个字符串列表，在原有slice上返回不包含空字符串的列表， 如：
	//	{"red", "", "black", "", "", "pink", "blue"}
	//——> {"red", "black", "pink", "blue"}
	/*
		slice_ :=[]string{"red", "", "black", "", "", "pink", "blue"}
		var slice_2 []string
		for _,v :=range slice_{
			if len(v) >0{
				fmt.Println(v)
				slice_2 = append(slice_2, v)
			}
		}
		fmt.Println(slice_2)*/
	//方法二：
	slice_ := []string{"red", "", "black", "", "", "pink", "blue"}
	i := 0
	for _, v := range slice_ {
		if v != "" {
			slice_[i] = v
			i++
		}
	}
	fmt.Println(slice_[:i]) //返回下标的切片位置

}

func slice_1_1() { //TODO 去重
	//练习2：写一个函数，就地消除[]string中重复字符串，如：
	//	{"red", "black", "red", "pink", "blue", "pink", "blue"}
	//	——>	{"red", "black", "pink", "blue"}
	data := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	out := data[:0] //创建一个空切片
	fmt.Println(out)
	for _, value := range data { //外层遍历要查询的数据
		i := 0
		for ; i < len(out); i++ { //内层需要定义一个循环来查看自己定义的空切片是否有重复的数据，如有就跳出循环如果没有记住遍历的位置以及下标的位置添加到空切片中
			if value == out[i] {
				break
			}

		}
		if i == len(out) {
			out = append(out, value)
		}
	}
	fmt.Println(out)

}

func remove() {
	slice_1a := []int{4, 5, 6, 7, 8, 9}
	//   目标位置切片：下标替换元素位置     源切片：下标确认元素
	copy(slice_1a[1:], slice_1a[3:])
	fmt.Println(slice_1a)
	fmt.Println(slice_1a[3:])
}

func wcFunc(str_text string) map[string]int {
	//练习3：	封装 wcFunc() 函数。接收一段英文字符串str。返回一个map，记录str中每个“词”出现次数的。
	//如："I love my work and I love my family too"
	//map[I:2 and:1 family:1 love:2 my:2 too:1 work:1]
	list_ := strings.Fields(str_text)
	m := map[string]int{}
	for i := 0; i < len(list_); i++ {
		if _, bool_vlaue := m[list_[i]]; bool_vlaue {
			m[list_[i]] += 1
		} else {
			m[list_[i]] = 1
		}
	}
	return m

}

//func main_01() {
//	a,b := 10,20
//	print("start",a,b)
//	swap(a,b) //值传递
//	swap_1(&a,&b)//值引用  &取地址
//	fmt.Println("stop_vluae",a,b)
//}

func main() {
	//slice_1()
	//noSame()
	//slice_1_1()
	//remove()
	str_text := "I love my work and I love my family too"
	a := wcFunc(str_text)
	fmt.Println(a)

}
