package chapter3_7

import (
	"log"
	"net/http"
	"net/url"
)

// x-www-form-urlencoded
// curl -d test=value http://www.baidu.com
func Init() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://www.baidu.com", values)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
