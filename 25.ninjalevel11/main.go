package main

import (
	"log"
	"fmt"
	"encoding/json"
	"time"
)

// #1
type person struct {
	Name string
	Age int
	Sayings []string
}

// #3
type customErr struct {
	name string
	timestamp time.Time
	severity int
}

func (c customErr) Error() string {
	return fmt.Sprintf("CustomErr with fields: %v %v %v", c.name, c.timestamp, c.severity)
}

func main() {
	// #1
	p1 := person{
		Name: "Andrew J",
		Age: 33,
		Sayings: []string{"hi", "hello", "wow"},
	}

	by, err := json.Marshal(p1)
	if err != nil {
		// fmt.Println(err)
		log.Fatalln("JSON did not marshal ,error: ", err)
	}
	fmt.Println(string(by))

	// #2
	s, errr := errorf("Afea")
	if errr != nil {
		fmt.Println(s, errr)
		// log.Println(errr)
	}
	fmt.Println("#2 done")

	// #3
	st, erro := customE("fewa")
	if erro != nil {
		fmt.Println(st, erro)
	}

	// #4
		// repeat of #3...can have error type field in a struct
	
	// #5 get testing up and running, see sample_test.go
}

// #2
func errorf(s string) (string, error) {
	if len(s) < 100 {
		return s, fmt.Errorf("string length is less than 100 bytes: %s", s)
		// return s, errors.New(fmt.Springf("error occurred with string: %s", s)) // need errors pkg
	}
	return s[50:], nil
}

// #3
func customE(s string) (string, error) {
	if s != "ORA ORA ORA" {
		return s, customErr{"string provided is not 'ORA ORA ORA'", time.Now(), 6}
	}
	return s[4:], nil
}

// #5
func returnString() string {
	return "Hello world!"
}