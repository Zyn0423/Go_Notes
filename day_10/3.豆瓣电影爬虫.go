package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

//URL ：https://movie.douban.com/top250

//https://movie.douban.com/top250?start=75&filter=

func err_fun2(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		return
	}
}

func Regxp_Compile_DouBan(re string, data []byte) []string {
	regx, err := regexp.Compile(re)
	err_fun2(err, "regexp.Compile")
	data_ := regx.FindAllStringSubmatch(string(data), -1)
	string_list := []string{}
	for _, v := range data_ {
		string_list = append(string_list, v[1])
	}
	return string_list

}

func parses_Html(html_data []byte) {
	Movie_Name := Regxp_Compile_DouBan(`<img width="100" alt="(.*?)"`, html_data)
	Movie_grade := Regxp_Compile_DouBan(`<span class="rating_num" property="v:average">(.*?)</span>`, html_data)
	Movie_evaluate := Regxp_Compile_DouBan(`<span>(.*?)人评价</span>`, html_data)
	for i := 0; i < 25; i++ {
		fmt.Println("电影名："+Movie_Name[i], "评分 ："+Movie_grade[i], "点赞人数 ："+Movie_evaluate[i])
	}

}
func Http_requests1(i int) (result string) {
	url_data := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
	//url_data :="https://tieba.baidu.com/f?kw=%E6%B5%B7%E8%B4%BC%E7%8E%8B&ie=utf-8&pn=0"
	//fmt.Printf("%s", url_data)

	client := &http.Client{}
	requests, err := http.NewRequest("GET", url_data, nil)
	requests.Header.Add("User-Agent", "Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	err_fun2(err, "http.NewRequest")
	respones, err := client.Do(requests)
	defer respones.Body.Close()
	err_fun2(err, "client.Do")
	html, err := ioutil.ReadAll(respones.Body)
	err_fun2(err, "ioutil.ReadAll")
	parses_Html(html)

	//fmt.Println(string(html))

	//resp ,err:=http.Get(url_data)
	//
	//if err != nil{
	//	fmt.Println("http.Get",err)
	//}
	//defer resp.Body.Close()
	//fmt.Println(resp.Header)
	//buf:=make([]byte,1024*4)
	//
	//fmt.Println("3")
	//for {
	//	n,err:=resp.Body.Read(buf)
	//	fmt.Println(string(buf[:n]))
	//	if n == 0 {
	//		fmt.Println("以结束")
	//		break
	//	}
	//	if err !=nil && err !=io.EOF{
	//		fmt.Println("io.EOF",err)
	//		return
	//	}
	//	result += string(buf[:n])
	//	//fmt.Println(result)
	//
	//}
	////return result

	return

}

func WorkIng(start, end int) {
	for i := start; i <= end; i++ {
		result := Http_requests1(i)
		fmt.Println(result)
		time.Sleep(time.Second * 10)
	}

}

func main() {

	var start, end int //TODO 定义起始和结束类型
	fmt.Print("输入起始页 ：")
	fmt.Scan(&start)
	fmt.Print("输入结束页 ：")
	fmt.Scan(&end)
	WorkIng(start, end)

}
