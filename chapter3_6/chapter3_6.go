package chapter3_6

import (
	"log"
	"net/http"
)

// curl --head http://www.baidu.com
func Init() {
	resp, err := http.Head("http://www.baidu.com")

	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
