package main

import (
	"fmt"
	"net"
	"os"
)

func Errfunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}
func main() {
	Conn, err := net.Dial("tcp", "127.0.0.1:8000")
	Errfunc(err, "net.Dial")
	defer Conn.Close()
	Conn.Write([]byte("GET /login HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"))

	buf := make([]byte, 1024*4)
	n, err := Conn.Read(buf)
	Errfunc(err, "Conn.Read")
	if n == 0 {
		return
	}
	fmt.Printf("|%s|\n", string(buf[:n]))
}
