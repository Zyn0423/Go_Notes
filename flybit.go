//package main
////
////import (
////	"fmt"
////	"io/ioutil"
////	"net/http"
////
////)
////
////func main() {
////
////	url := "https://api.flybit.com/v2/private/walletInfo"
////
////	req, _ := http.NewRequest("POST", url, nil)
////
////	req.Header.Add("content-type", "application/json")
////	req.Header.Add("authorization", ".eyJhY2Nlc3NLZXkiOiJNakF5TUMwd05DMHdNbFF4TVRveU50WkRjd1pHWmpPR1V0TW1FMVlpMDBOREUwTFRneU1qRXRNMkU1Tnpsa01qTmhaVFUyTVRRME5UTXgiLCJub25jZGNhc2QxMmVhczEycTJlMSJ9.7QdOEXA5NC6pFI78LfYD-Hdg6HlCpZPk1Pz")
////	req.Header.Add("cache-control", "no-cache")
////	req.Header.Add("postman-token", "6bbe7ffd-d18b-3601-0506-bbfe5983dc3c")
////
////	res, _ := http.DefaultClient.Do(req)
////
////	defer res.Body.Close()
////	body, _ := ioutil.ReadAll(res.Body)
////
////	fmt.Println("res",res)
////	fmt.Println(string(body))
////
////}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://api.flybit.com/v2/private/order"

	//payload := strings.NewReader("{\r\n    \"marketId\": 1,\r\n    \"amount\": 1,\r\n    \"orderType\": \"B\",\r\n    \"callType\": \"N\",\r\n    \"price\": 8000000\r\n}")
	//payload := strings.NewReader("{\n    \"marketId\": 1,\n    \"amount\": 1,\n    \"orderType\": \"B\",\n    \"callType\": \"N\",\n    \"price\": 8000000\n} 0 -1}\n")
	payload := strings.NewReader("{\n    \"marketId\": 1,\n    \"amount\": 1,\n    \"orderType\": \"B\",\n    \"callType\": \"N\",\n    \"price\": 8000000\n}")
	fmt.Println(*payload)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", ""+
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhY2Nlc3NLZXkiOiJNakF5TUMwd05DMHdObFF4TnpveE9Eb3pOMW90Wm1abU9ETXhZakl0WW1VeU5pMDBaVEF6TFdKbVlUY3RNalZqWm1WbU1UUmpZamhrTVRRME5UTXgiLCJub25jZSI6IjIwMjAtMDQtMDhUMTE6MTE6NTlaYTMwYzY4NTgyZGFhNGFmMTlmNzI4ZTI1MjIzMmQyZDMifQ.202EWPgKoxmjfuc9SvjLauqSlA86CsE6PnVCnZ8Jn6I")
	fmt.Println(req.Header)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))

}
