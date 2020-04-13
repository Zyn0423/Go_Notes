package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	buf := make([]byte, 1024*4)
	var string_data string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("数据读完毕", err)
			break
		}
		if err != nil && io.EOF == nil {
			fmt.Println("resp.Body.Read", err)
			break

		}
		string_data += string(buf[:n])
	}
	fmt.Printf("%v", string_data)

}
