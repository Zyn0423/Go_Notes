package main

import (
	"fmt"
	"os"
	"strings"
)

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
				fmt.Println("jpg 文件有：", fileInfo.Name())
			}
		}
	}
}
