package chapter3_19

import (
	"log"
	"net/http"
	"net/http/httputil"
)

// curl -X DELETE http://www.baidu.com
func Init() {
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	log.Println(string(dump))
}
