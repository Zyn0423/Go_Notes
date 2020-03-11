package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//	先打开文件，然后在创建一份文件，讲打开的文件边读边写
	f_r, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer f_r.Close()

	f_w, err := os.Create("./test1.txt")

	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer f_w.Close()
	//TODO 从读文件当中放入缓冲区中
	buf := make([]byte, 1024*4)
	//TODO 循环从读文件中，获取数据，“原封不动的”写到写文件中。
	for {
		n, err := f_r.Read(buf)
		if err != nil && io.EOF == err {
			fmt.Printf("写入完毕 %d\n", n)
			return

		}
		f_w.Write(buf[:n]) // TODO 读多少，写多少

	}

}
