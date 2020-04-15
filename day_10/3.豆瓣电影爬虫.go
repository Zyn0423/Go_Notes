//package main
//
//import (
//	"fmt"
//	"io"
//	"net/http"
//	"strconv"
//	"time"
//)
//
////URL ：https://movie.douban.com/top250
//
////https://movie.douban.com/top250?start=75&filter=
//
//func Http_requests1(i int)(result string)  {
//	url_data :="https://movie.douban.com/top250?start="+strconv.Itoa((i-1)*25)+"&filter="
//	//url_data :="https://tieba.baidu.com/f?kw=%E6%B5%B7%E8%B4%BC%E7%8E%8B&ie=utf-8&pn=0"
//	fmt.Printf("%s",url_data)
//
//
//	resp ,err:=http.Get(url_data)
//
//	if err != nil{
//		fmt.Println("http.Get",err)
//	}
//	defer resp.Body.Close()
//	fmt.Println(resp.Header)
//	buf:=make([]byte,1024*4)
//
//	fmt.Println("3")
//	for {
//		n,err:=resp.Body.Read(buf)
//		fmt.Println(string(buf[:n]))
//		if n == 0 {
//			fmt.Println("以结束")
//			break
//		}
//		if err !=nil && err !=io.EOF{
//			fmt.Println("io.EOF",err)
//			return
//		}
//		result += string(buf[:n])
//		//fmt.Println(result)
//
//	}
//	//return result
//
//	return
//
//}
//
//func WorkIng(start,end int)  {
//	for i:=start ;i<=end;i++{
//		result:=Http_requests1(i)
//		fmt.Println(result)
//		time.Sleep(time.Second*10)
//	}
//
//
//}
//
//
//func main() {
//	var start, end int   //TODO 定义起始和结束类型
//	fmt.Print("输入起始页 ：")
//	fmt.Scan(&start)
//	fmt.Print("输入结束页 ：")
//	fmt.Scan(&end)
//	WorkIng(start,end)
//
//
//}
//

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func work_ing3(url_ string) {
	clt := &http.Client{}
	request, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		fmt.Println("http.NewRequest", err)
	}
	request.Header.Add("User-Agent", "Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

	Resp, err := clt.Do(request)
	if err != nil {
		fmt.Println("clt.Do", err)
	}
	defer Resp.Body.Close()
	data, err := ioutil.ReadAll(Resp.Body)
	fmt.Println(Resp)
	if err != nil {
		fmt.Println("ioutil.ReadAll", err)
	}
	fmt.Println(string(data))

}
func main() {
	url_baidu := "https://www.baidu.com/s?wd=汇率"
	work_ing3(url_baidu)
}
