package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type bankerrate struct {
	Name             string
	Spotexchangebuy  string
	Oofbuy           string
	Spotexchangesell string
	Oofsell          string
	Middleprice      string
	Issuetime        string
}

func Map_bobyplant(boby map[string]string) (string_boby string) { //TODO boby 参数加工厂
	string_boby = ""
	for k, v := range boby {

		string_boby += k + "=" + v + ","
	}
	return
}

func err_fun(err error, info string) { // todo 错误函数
	if err != nil {
		fmt.Println(info, err)
		return
	}

}

func list_data(data *bankerrate, list_split []string) {
	data.Name = list_split[1]
	data.Spotexchangebuy = list_split[2]
	data.Oofbuy = list_split[3]
	data.Spotexchangesell = list_split[4]
	data.Oofsell = list_split[5]
	data.Middleprice = list_split[6]
	data.Issuetime = list_split[7]
	fmt.Println(*data)

}
func parase_html(data string) { //TODO 解析函数
	rex, err := regexp.Compile(`<tr>(?s:(.*?))</tr>`)
	err_fun(err, "regexp.Compile")                // 正则语法
	result := rex.FindAllStringSubmatch(data, -1) // 获取筛选后的数据
	var b bankerrate                              //声明b 变量结构体

	for _, n := range result {
		rex, err := regexp.Compile(`<td>(?s:(.*?))</td>`) // 正则语法
		err_fun(err, "regexp.Compile")
		result2 := rex.FindAllStringSubmatch(n[1], -1) // 获取筛选后的数据
		string_data := ""                              // TODO 对数据进行拼接
		for _, n1 := range result2 {
			string_data += " " + n1[1]
		}

		strlist := strings.Split(string_data, " ")
		if len(strlist) > 1 {
			list_data(&b, strlist)
		} else {
			continue
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
