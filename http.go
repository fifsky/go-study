package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/goapt/golib/httpclient"
)

func main() {
	http := httpclient.NewClient()

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	str, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(str))

	var p = url.Values{}
	p.Add("app_id", "ccccc")
	p.Add("app_key", "dddddddddddddd")
	body := strings.NewReader(p.Encode())

	httpResp, err := http.Post("https://httpbin.org/post", "application/xml; charset=utf-8", body)

	if err != nil {
		fmt.Println(err)
	}
	defer httpResp.Body.Close()

	str, _ = ioutil.ReadAll(httpResp.Body)
	fmt.Println(string(str))
}
