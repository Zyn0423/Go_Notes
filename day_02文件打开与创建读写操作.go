package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func create_file() { //TODO 新建文件，文件不存在创建，文件存在将清空

	f, err := os.Create("./test.txt") //接收两个参数
	if err != nil {
		fmt.Println("创建失败", err)
		return
	}

	defer f.Close()
	fmt.Println("successful")

}

func open_file() { //TODO 打开文件并写入 以只读只写方式打开
	/*
		os.O_RDONLY: 以只读的方式打开
		os.O_WRONLY: 以只写的方式打开
		os.O_RDWR : 以读写的方式打开
		os.O_NONBLOCK: 打开时不阻塞
		os.O_APPEND: 以追加的方式打开
		os.O_CREAT: 创建并打开一个新文件
		os.O_TRUNC: 打开一个文件并截断它的长度为零（必须有写权限）
		os.O_EXCL: 如果指定的文件存在，返回错误
		os.O_SHLOCK: 自动获取共享锁
		os.O_EXLOCK: 自动获取独立锁
		os.O_DIRECT: 消除或减少缓存效果
		os.O_FSYNC : 同步写入
		os.O_NOFOLLOW: 不追踪软链接
	*/
	//flag：
	//perm : 一般有7中权限默认传6
	f, err := os.OpenFile("./test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("打开失败", err)
		return

	}
	defer f.Close()
	_, err = f.WriteString("Hello world !!!!\r") //TODO 直接换行\r
	if err != nil {
		fmt.Printf("写入失败", err)
		return
	}
	fmt.Println("successful")
	off, _ := f.Seek(-5, io.SeekEnd) //TODO  修改文件的读写指针的位置  offset int64（偏移量  正：向文件尾偏， 负：向文件头偏）, whence int（起始位置）  返回值ret代表偏移量：从起始位置的值
	//	SeekStart   = 0 // seek relative to the origin of the file  文件起始位置
	//	SeekCurrent = 1 // seek relative to the current offset     文件当前位置
	//	SeekEnd     = 2 // seek relative to the end              文件结尾位置

	n, _ := f.WriteAt([]byte("1111"), off) //TODO 文件制定偏移位置，写入[]byte,通常搭配Seek()使用 参数：1，写入的数据2.偏移量  返回：实际写出的偏移量
	fmt.Printf("d", n)
}
func read_file() { //TODO 读文件操作

	f, err := os.OpenFile("./test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("打开失败", err)
		return

	}
	defer f.Close()
	//TODO 创建一个缓冲区读写器
	reader := bufio.NewReader(f) //TODO 按行读打开文件和指针   返回的red类型
	for {
		buf, err := reader.ReadBytes('\r') //TODO 传入字节\n  从reader的缓冲区中，读取指定长度的数据，数据长度取决于参数的delim
		if err != nil && err == io.EOF {
			//TODO io.EOF 文件结束标记，是要单独读一次获取到
			fmt.Println("读取文件完毕")
			return
		} else if err != nil {
			fmt.Println("ReadBytes 失败", err)
		}
		fmt.Print(string(buf))
	}

}

func main() {
	//create_file()
	//open_file()
	read_file()
}
