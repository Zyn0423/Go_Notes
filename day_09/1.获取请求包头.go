package main

import (
	"fmt"
	"net"
	"os"
)

func errfunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}
func main() {
	Listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errfunc(err, "net.Listen")
	defer Listener.Close()
	Conn, err := Listener.Accept()
	errfunc(err, "Listener.Accept")
	defer Conn.Close()
	buf := make([]byte, 1024*4)
	n, err := Conn.Read(buf)
	errfunc(err, "Conn.Read")
	fmt.Printf("|%s|", string(buf[:n]))

	//	|GET / HTTP/1.1
	//Host: 127.0.0.1:8000
	//Connection: keep-alive
	//Upgrade-Insecure-Requests: 1
	//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36
	//Sec-Fetch-Dest: document
	//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
	//Sec-Fetch-Site: none
	//Sec-Fetch-Mode: navigate
	//Accept-Encoding: gzip, deflate, br
	//Accept-Language: zh-CN,zh;q=0.9
	//Cookie: go_admin_session=41e4f39c-9987-4aca-ac50; Hm_lvt_d97abf6d61
	//
	//|
}
