package main

import (
	"fmt"
	"net/http"
	"os"
)

func opensendfile(filename string, w http.ResponseWriter) {
	FileNamePath := "/Users/xxx/xxxx" + filename
	f, err := os.Open(FileNamePath)

	if err != nil {
		fmt.Println("os.Open :", err)
		w.Write([]byte("No such file or directory"))
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*4)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			return
		}
		w.Write(buf[:n])
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客服端请求", r.URL)
	opensendfile(r.URL.String(), w)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
