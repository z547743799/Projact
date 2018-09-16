package main

import (
	"github.com/ddliu/go-httpclient"
)

const (
	USERAGENT       = "my awsome httpclient"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 5
	SERVER          = "https://github.com"
)

func main() {
	httpclient.Defaults(httpclient.Map{
		"opt_useragent":   USERAGENT,
		"opt_timeout":     TIMEOUT,
		"Accept-Encoding": "gzip,deflate,sdch",
	})

	//res, _ := httpclient.Get("http://www.baidu.com")
	//
	//
	//fmt.Println("Response:")
	//fmt.Println(res.ToString())

	//rs,_:=httpclient.Get("http://httpbin.org/get", map[string]string{
	//	"q": "news",
	//})
  //fmt.Println(rs.ToString())


}