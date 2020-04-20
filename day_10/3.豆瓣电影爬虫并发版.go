package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

//URL ：https://movie.douban.com/top250

//https://movie.douban.com/top250?start=75&filter=

func err_fun3(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		return
	}
}

func Regxp_Compile_DouBan1(re string, data []byte) []string {
	regx, err := regexp.Compile(re)
	err_fun3(err, "regexp.Compile")
	data_ := regx.FindAllStringSubmatch(string(data), -1)
	string_list := []string{}
	for _, v := range data_ {
		string_list = append(string_list, v[1])
	}
	return string_list

}

func parses_Html1(html_data []byte) {
	Movie_Name := Regxp_Compile_DouBan1(`<img width="100" alt="(.*?)"`, html_data)
	Movie_grade := Regxp_Compile_DouBan1(`<span class="rating_num" property="v:average">(.*?)</span>`, html_data)
	Movie_evaluate := Regxp_Compile_DouBan1(`<span>(.*?)人评价</span>`, html_data)
	for i := 0; i < 25; i++ {
		fmt.Println("页 电影名："+Movie_Name[i], "评分 ："+Movie_grade[i], "点赞人数 ："+Movie_evaluate[i])
	}

}
func Http_requests2(i int, page chan int) (result string) {
	url_data := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
	client := &http.Client{}
	requests, err := http.NewRequest("GET", url_data, nil)
	requests.Header.Add("User-Agent", "Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	err_fun3(err, "http.NewRequest")
	respones, err := client.Do(requests)
	defer respones.Body.Close()
	err_fun3(err, "client.Do")
	html, err := ioutil.ReadAll(respones.Body)
	err_fun3(err, "ioutil.ReadAll")
	parses_Html1(html)
	page <- i
	return

}

func WorkIng1(start, end int) {
	page := make(chan int)
	for i := start; i <= end; i++ {
		go Http_requests2(i, page)
		//time.Sleep(time.Second * 10)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d 页爬取完毕\n", <-page)
	}

}

func main() {

	var start, end int //TODO 定义起始和结束类型
	fmt.Print("输入起始页 ：")
	fmt.Scan(&start)
	fmt.Print("输入结束页 ：")
	fmt.Scan(&end)
	WorkIng1(start, end)

}
