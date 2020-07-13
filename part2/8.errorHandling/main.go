package main

import (
	"fmt"
	"time"
)

type cuserr struct {
	name string
	time time.Time
	severity int
}

func (c cuserr) Error() string {
	return fmt.Sprintf("name: %s, time: %v, severity: %d", c.name, c.time, c.severity)
}

func main() {
	val, err := customE()
	if err != nil {
		fmt.Println(val, err)
	}
}

func customE() (string, error) {
	return "finished", cuserr {
		name: "custome error",
		time: time.Now(),
		severity: 5,
	}
}