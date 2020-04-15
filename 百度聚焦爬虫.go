package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 编码
//func ConvertToString(src string, srcCode string, tagCode string) string {
//	srcCoder := mahonia.NewDecoder(srcCode)
//	srcResult := srcCoder.ConvertString(src)
//	tagCoder := mahonia.NewDecoder(tagCode)
//	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
//	result := string(cdata)
//	return result
//}

func main() {
	url := "https://www.baidu.com/s?wd=%E6%B1%87%E7%8E%87"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Cookie", "BIDUPSID=4B1603EE0F08DB4DBFE3093B91AEA24E; PSTM=1586921512; BAIDUID=4B1603EE0F08DB4DAE6EAC0B5FD2BE24:FG=1; delPer=0; BD_CK_SAM=1; PSINO=7; H_PS_PSSID=30962_1428_31125_21114_31187_31270_31051_30823_26350_31163_31196; BDSVRTM=10")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	res, err := client.Do(req)
	//req.Header.Set("Accept-Encoding", "identity")
	defer res.Body.Close()
	//data ,err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))

	if res.StatusCode == 200 {
		h := res.Header
		for k, v := range h {
			fmt.Println(k, v)
		}
		fmt.Println(res.Header.Get("Content-Encoding"))

		reader, _ := gzip.NewReader(res.Body)

		for {
			buf := make([]byte, 1024*4)
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF && !strings.Contains(err.Error(), "EOF") {
				return
			}
			if n == 0 {
				return
			}
			fmt.Println(string(buf))
		}
	} else {
		fmt.Println(res.StatusCode)
	}

	//body ,err :=gzip.NewReader(res.Body)
	//if err !=nil{
	//	fmt.Println("gzip.NewReader",err)
	//}
	//defer body.Close()
	//
	//data ,err := ioutil.ReadAll(body)

	//fmt.Println(string(data))

	//op_exrate_sub_title

	//var reader io.ReadCloser
	//switch res.Header.Get("Content-Encoding") {
	//case "gzip":
	//	reader, err = gzip.NewReader(res.Body)
	//	defer reader.Close()
	//default:
	//	reader = res.Body
	//}
	//fmt.Println(reader)

	//_,_=io.Copy(os.Stdout, reader) // print html to standard out
	//fmt.Printf("%T",string(n))

	/*
		buf :=make([]byte,1024*4)
		for {
			n ,err := res.Body.Read(buf)

			if n == 0 {
				fmt.Println("读取网页完成")
				break
			}
			if err!=nil && err!=io.EOF{
				fmt.Println("res.Body.Read",err)
			}
			//result:=strconv.Itoa(n)
			//result:=string(buf[:n])


			//ParseGzip(buf[:n])






			//解决中文乱码
			//bodystr := mahonia.NewDecoder("utf8").ConvertString(string(buf[:n]))
			//fmt.Printf("%v",bodystr)
			//fmt.Println(bodystr)

		}
	*/

	//
	//body, err := ioutil.ReadAll(res.Body)
	//fmt.Printf("%T\n",body)
	//var a uint strconv.Itoa(int(a))
	//fmt.Printf("%v",body)

	//fmt.Println(string(body))
}
