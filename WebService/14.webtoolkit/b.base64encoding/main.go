package main

import(
	"fmt"
	"encoding/base64"
)

func main() {
	s := "some string"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkl{}[]`~_+><.,:;!@#$%^&*()/?"

	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	ss64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s), len(s64), len(ss64),  s, s64, ss64)

	sd, err := base64.NewEncoding(encodeStd).DecodeString(s64)
	if err != nil {
		fmt.Println(err)
	}

	ssd, err := base64.StdEncoding.DecodeString(ss64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(sd), len(ssd), string(sd), string(ssd))
}