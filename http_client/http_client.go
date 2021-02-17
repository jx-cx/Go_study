package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//resp, err := http.Get("http://127.0.0.1:9000/xxx/?name=sb&age=18")
	//if err != nil {
	//	fmt.Printf("get url failed,err:%v", err)
	//	return
	//}
	data := url.Values{} // URL values
	urlObj, _ := url.Parse("http://127.0.0.1:9000/xxx/")
	data.Set("name", "傻逼")
	data.Set("age", "88")
	queryStr := data.Encode() //URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	fmt.Printf("get url failed,err:%v\n", err)
	//}
	//禁用keepAlive的client
	tr := &http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get url failed,err:%v\n", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Sprintf("read resp.body failed,err:%v", err)
		return
	}
	fmt.Println(string(b))
}
