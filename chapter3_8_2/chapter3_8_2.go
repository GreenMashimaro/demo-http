package chapter3_8_2

import (
	"log"
	"net/http"
	"strings"
)

func Init() {
	reader := strings.NewReader("hello world")
	resp, err := http.Post("http://www.baidu.com", "text/plain", reader)

	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
