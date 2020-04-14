package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func err_fun1(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}

func Http_requests(url_baidu string) (result string) {
	resp, err := http.Get(url_baidu)
	err_fun1(err, "http.Get")
	defer resp.Body.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			return
		}
		if err != nil && err != io.EOF {
			return
		}
		// 累加每一次循环读到的 buf 数据，存入result 一次性返回。
		result += string(buf[:n])
	}
}

func spider(i int, page chan int) {
	url_baidu := "https://tieba.baidu.com/f?kw=%E6%B5%B7%E8%B4%BC%E7%8E%8B&ie=utf-8&pn="

	result := Http_requests(url_baidu + strconv.Itoa((i-1)*50))

	f, err := os.Create("第 " + strconv.Itoa(i) + " 页" + ".html")
	if err != nil {
		fmt.Println("os.Create", err)
		return
	}
	f.WriteString(result)
	f.Close()

	page <- i // 与主go程完成同步。

}

func work_ing2(start, end int) {
	page := make(chan int)
	// 循环爬取每一页数据
	for i := start; i <= end; i++ {
		go spider(i, page)

	}
	for i := start; i <= end; i++ {
		fmt.Println("已完成 ：", <-page)

	}

}
func main() {
	// 指定爬取起始、终止页
	var start, end int
	fmt.Print("请输入起始页 ：")
	fmt.Scan(&start)
	fmt.Print("请输入结束页 ：")
	fmt.Scan(&end)
	work_ing2(start, end)

}
