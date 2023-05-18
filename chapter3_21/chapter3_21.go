package chapter3_21

import (
	"fmt"
	"golang.org/x/net/idna"
)

func Init() {
	src := "握力王"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s -> %s\n", src, ascii)
}
