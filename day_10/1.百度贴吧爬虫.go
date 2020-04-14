package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func err_fun(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}

}

func err_fun_file(err error, info string) {
	if err != nil && io.EOF != nil {
		fmt.Println(info, err)
		return
	}
}
func HttpRespone(url_baidu string) (result string) {
	resp, err := http.Get(url_baidu) //Itoa是FormatInt(i, 10) 的简写。
	defer resp.Body.Close()
	err_fun(err, "http.Get")
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		err_fun_file(err, "resp.Body.Read")
		//fmt.Println(string(buf[:n]))
		result += string(buf[:n])
	}
	return

}

func work_ing1(start, end int, url_baidu string) {
	fmt.Println("start", start)
	fmt.Println("end", end)
	for i := start; i <= end; i++ {

		result := HttpRespone(url_baidu + strconv.Itoa((i-1)*50))
		fmt.Println("已爬取 ：", (i-1)*50)
		fmt.Println(result)
		f, err := os.Create("第 " + strconv.Itoa(i) + " 页" + ".html")
		if err != nil {
			fmt.Println("Create err:", err)
			continue
		}
		f.WriteString(result)
		f.Close()

	}

}

func main() {
	var start, end int
	fmt.Print("请输入起始页 ：")
	fmt.Scan(&start)
	fmt.Print("请输入结束页 ：")
	fmt.Scan(&end)
	url_baidu := "https://tieba.baidu.com/f?kw=%E6%B5%B7%E8%B4%BC%E7%8E%8B&ie=utf-8&pn="
	work_ing1(start, end, url_baidu)

}
