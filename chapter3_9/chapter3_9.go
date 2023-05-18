package chapter3_9

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// multipart/form-data
// curl -F "name=Michael Jackson" -F "thumbnail=@photo.png" http://www.baidu.com
func Init() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	fileWrite, err := writer.CreateFormFile("thumbnail", "photo.png")

	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("./chapter3_9/photo.png")

	// open file fail
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	io.Copy(fileWrite, readFile)

	// Content-Type: application/octet-stream
	resp, err := http.Post("http://www.baidu.com", writer.FormDataContentType(), &buffer)

	// send fail
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
