package main

import (
	"fmt"
	"io"
	"crypto/sha256"
	"crypto/hmac"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@example.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}