package main

import (
	"fmt"
)

func demo01() {
	fmt.Println("hello world")
}
func main_01() { //comma -ok断言
	var p []interface{}
	p = append(p, 123, "dada", 3.14, demo01)
	for _, v := range p {
		if value, ok := v.(int); ok {
			fmt.Println("整型类型", value)
		} else if value, ok := v.(string); ok {
			fmt.Println("字符串类型", value)
		} else if value, ok := v.(func()); ok {
			value()
			fmt.Println("函数类型", value)
		} else if value, ok := v.(float64); ok {
			fmt.Println("浮点类型", value)
		}

	}
}

func type_custom(a []interface{}) { //接收参数类型是切片接口类型
	//fmt.Println(a)
	for _, v := range a {
		switch value := v.(type) {
		case int:
			fmt.Println("整型类型", value)
		case string:
			fmt.Println("字符串类型", value)
		case float64:
			fmt.Println("浮点类型", value)
		case func():
			value()
			fmt.Println("浮点类型", value)
		}

	}

}

func main_02() { //switch
	var p []interface{}
	p = append(p, 123, "dada", 3.14, demo01)
	for _, v := range p {
		switch value := v.(type) {
		case int:
			fmt.Println("整型类型", value)
		case string:
			fmt.Println("字符串类型", value)
		case float64:
			fmt.Println("浮点型", value)
		case func():
			value()
			fmt.Println("函数类型", value)

		}
	}
}

func main() {
	var p []interface{}

	p = append(p, 123, "dada", 3.14, demo01)
	type_custom(p)
}
