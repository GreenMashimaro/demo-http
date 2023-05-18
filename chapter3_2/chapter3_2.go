package chapter3_2

import (
	"io/ioutil"
	"log"
	"net/http"
)

// curl https://www.baidu.com
func Init() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
}
