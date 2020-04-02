package main

import (
	"fmt"
	"os"
)

func main() {
	list_args := os.Args //TODO 获取命令行数据
	//fmt.Println(list_args[3])
	if len(list_args) != 2 {
		fmt.Println("请输入要路径")
	}

	fileinfo, err := os.Stat(list_args[1])
	if err != nil {
		fmt.Println("os.Stat", fileinfo)
	}
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	//Name() string       // base name of the file
	//Size() int64        // length in bytes for regular files; system-dependent for others
	//Mode() FileMode     // file mode bits
	//ModTime() time.Time // modification time
	//IsDir() bool        // abbreviation for Mode().IsDir()
	//Sys() interface{}   // underlying data source (can return nil)

}
