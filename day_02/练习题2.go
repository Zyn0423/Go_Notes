package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 拷贝mp3 文件到指定目录的操作
func cpMP32Dir(src, dst string) {
	//fmt.Println("src:", src)
	//fmt.Println("dst:", dst)

	//打开读文件
	f_r, err := os.Open(src)
	if err != nil {
		fmt.Println("Open err: ", err)
		return
	}
	defer f_r.Close()
	// 创建写文件
	f_w, err := os.Create(dst)
	if err != nil {
		fmt.Println("Create err: ", err)
		return
	}
	defer f_w.Close()

	// 从读文件中获取数据，放到缓冲区中。
	buf := make([]byte, 4096)
	// 循环从读文件中，获取数据，“原封不动的”写到写文件中。
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("读完。n = %d\n", n)
			return
		}
		f_w.Write(buf[:n]) // 读多少，写多少
	}
}

func main() {
	// 获取用户输入的目录路径
	fmt.Println("请输入待查询的目录：")
	var path string
	fmt.Scan(&path)

	// 打开目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err: ", err)
		return
	}
	defer f.Close()

	// 读取目录项
	info, err := f.Readdir(-1) // -1： 读取目录中所有目录项
	if err != nil {
		fmt.Println("Readdir err: ", err)
		return
	}
	// 变量返回的切片
	for _, fileInfo := range info {
		if !fileInfo.IsDir() { // 文件
			if strings.HasSuffix(fileInfo.Name(), ".mp3") {
				cpMP32Dir(path+"/"+fileInfo.Name(), "./"+fileInfo.Name())
			}
		}
	}
}
