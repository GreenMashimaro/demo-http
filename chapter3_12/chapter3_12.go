package chapter3_12

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
)

func Init() {
	jar, err := cookiejar.New(nil)

	if err != nil {
		panic(err)
	}

	client := http.Client{
		Jar: jar,
	}

	// 由于Cookie的结构是客户端第一次访问服务器时接收Cookie，第二次及之后访问时向服务器发送Cookie，因为这里进行两次访问
	for i := 0; i < 2; i++ {
		resp, err := client.Get("http://www.baidu.com")
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}

		log.Println("dump.length", len(dump))
		log.Println("Status:", resp.Status)
	}
}
