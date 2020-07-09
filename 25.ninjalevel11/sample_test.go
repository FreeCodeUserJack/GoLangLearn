package main

import (
	"fmt"
	"testing"
)

func TestReturnString(t *testing.T) {
	res := returnString()
	if res != "Hello world!" {
		fmt.Println("res doesn't match 'Hello world!'")
		t.Errorf("%s doesn't match 'Hello world!'", res)
	} else {
		fmt.Println("returnString returned string is equal to  'Hello world!'")
	}
}