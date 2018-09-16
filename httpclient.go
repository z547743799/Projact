package main

func main() {
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "my awsome httpclient",
		"Accept-Language":        "en-us",
	})

	res, err := httpclient.Get("http://google.com/search", map[string]string{
		"q": "news",
	})

	println(res.StatusCode, err)
}
