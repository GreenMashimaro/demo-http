package chapter3_4

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// curl --get --data-urlencode "query=hello world" http://www.baidu.com
func Init() {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get("http://www.baidu.com" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
