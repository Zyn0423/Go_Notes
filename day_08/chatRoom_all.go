package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client3 struct {
	C     chan string
	Name  string
	Adder string
}

var messagesclientmap3 map[string]Client3

var messages3 = make(chan string)

func WriteMsgToClient3(conn net.Conn, clt Client3) {
	for msg := range clt.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func MakeMsg3(clt Client3, data string) (buf string) {
	buf = "[" + clt.Adder + "] :" + clt.Name + " :" + data
	return
}
func handlerConnet3(conn net.Conn) {
	defer conn.Close()
	hasdata := make(chan bool) // 监听用户是否在线
	Addr := conn.RemoteAddr().String()
	clt := Client3{make(chan string), Addr, Addr}
	go WriteMsgToClient3(conn, clt)
	messagesclientmap3[Addr] = clt

	messages3 <- MakeMsg3(clt, "login")
	isQuit := make(chan bool)
	go func() {
		buf := make([]byte, 1024*4)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Printf("检测到用户端%s退出\n", clt.Name)
				isQuit <- true
				return
			}
			if err != nil {
				fmt.Println("conn.Write", err)
				return
			}

			//将读到用户消息保存到msg中
			msg := string(buf[:n-1])
			if msg == "who" && len(msg) == 3 {
				for _, value := range messagesclientmap3 {
					msg_list := "用户列表[" + value.Name + value.Adder + "]\n"
					conn.Write([]byte(msg_list))
				}
				//rename
			} else if len(msg) > 8 && msg[:6] == "rename" { // rename|
				NewName := strings.Split(msg, "|")[1] // msg[8:]
				clt.Name = NewName                    // 修改结构体成员name
				messagesclientmap3[Addr] = clt        // 更新 onlineMap
				conn.Write([]byte("rename successful\n"))
			} else {
				messages3 <- MakeMsg3(clt, msg)
			}
			hasdata <- true

		}

	}()
	for {
		select {
		case <-isQuit:
			delete(messagesclientmap3, clt.Adder) //将用户从messagesclientmap3 中移除
			messages3 <- MakeMsg3(clt, "exit")    //写入用户退出消息的全局channel
			return
		case <-hasdata: //什么都不做目的就是重置下面的定时器

		case <-time.After(time.Second * 10):
			delete(messagesclientmap3, clt.Adder) //将用户从messagesclientmap3 中移除
			messages3 <- MakeMsg3(clt, "exit")    //写入用户退出消息的全局channel
			return
		}
	}

}
func Manager3() {
	messagesclientmap3 = make(map[string]Client3)
	for {
		msg := <-messages3
		for _, clent := range messagesclientmap3 {
			clent.C <- msg
		}
	}
}
func main() {
	Listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Printf("net.Listen", err)
		return
	}

	defer Listener.Close()
	go Manager3()
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			fmt.Printf("Listener.Accept", err)
			return
		}
		go handlerConnet3(Conn)
	}
}
