package chapter3_15

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// curl -x http://www.baidu.com http://github.com
func Init() {
	proxyUrl, err := url.Parse("http://www.baidu.com")

	if err != nil {
		panic(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	resp, err := client.Get("http://github.com")
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
}
