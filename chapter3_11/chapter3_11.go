package chapter3_11

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func Init() {
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/png")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.png"`)

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")
	fileWrite, err := writer.CreatePart(part)

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
