package chapter3_8

import (
	"log"
	"net/http"
	"os"
)

// multipart 表单
// curl -T data.txt -H "Content-Type: text/plain" http://www.baidu.com
func Init() {
	file, err := os.Open("./chapter3_8/data.txt")

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://www.baidu.com", "text/plain", file)

	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
