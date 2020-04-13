package main

import (
	"fmt"
	"net/http"
)

//func handler(w http.ResponseWriter, r *http.Request)  {
//	w.Write([]byte("hello world"))
//	fmt.Println("Header",r.Header)
//	fmt.Println("Host",r.Host)
//	fmt.Println("RemoteAddr",r.RemoteAddr)
//	fmt.Println("Method",r.Method)
//	fmt.Println("Body",r.Body)
//	fmt.Println(r.URL)
//}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8001", nil)
}
