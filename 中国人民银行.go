package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Map_bobyplant(boby map[string]string) (string_boby string) {
	string_boby = ""
	for k, v := range boby {

		string_boby += k + "=" + v + ","
	}
	return
}

func err_fun(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		return
	}

}
func parase_html(data string) {
	rex, err := regexp.Compile(`<tr>(?s:(.*?))</tr>`)
	err_fun(err, "regexp.Compile")
	result := rex.FindAllStringSubmatch(data, -1)
	for _, n := range result {
		rex, err := regexp.Compile(`<td>(?s:(.*?))</td>`)
		err_fun(err, "regexp.Compile")
		result2 := rex.FindAllStringSubmatch(n[1], -1)
		for _, n1 := range result2 {
			fmt.Println(n1[1])
		}
	}
}

func main() {
	Client := &http.Client{}
	method := "POST"
	url := "https://srh.bankofchina.com/search/whpj/search_cn.jsp?pjname=美元&head=head_620.js&bottom=bottom_591.js"
	//boby := make(map[string]string)
	//boby["pjname"] = "美元"
	//boby["head"] = "head_620.js"
	//boby["bottom"] = "bottom_591.js"
	//string_boby:=Map_bobyplant(boby)
	//boby_str := strings.TrimRight(string_boby, ",") //字符串去除最后一个字符方法
	//fmt.Println("boby_str",boby_str)
	//Request, err := http.NewRequest(method, url, strings.NewReader(boby_str))
	Request, err := http.NewRequest(method, url, nil)
	err_fun(err, "http.NewRequest") //   错误函数
	Request.Header.Add("User-Agent", "Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

	respones, err := Client.Do(Request)
	err_fun(err, "http.NewRequest") //   错误函数
	n, err := ioutil.ReadAll(respones.Body)
	err_fun(err, "http.NewRequest") //   错误函数
	parase_html(string(n))
}
