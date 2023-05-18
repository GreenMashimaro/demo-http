package chapter3_20

import (
	"context"
	"log"
	"net/http"
	"time"
)

// curl -m 2 http://www.baidu.com/slow_page
func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", "http://www.baidu.com/slow_page", nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	log.Println("Status", res.Status)
}
