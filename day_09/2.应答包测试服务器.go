package main

import "net/http"

func handler1(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello world"))

}
func main() {
	http.HandleFunc("/login", handler1)        //todo 注册回调函数，该回调函数在服务器被访问时，自动会被调用
	http.ListenAndServe("127.0.0.1:8000", nil) //todo 绑定服务器监听地址

}
