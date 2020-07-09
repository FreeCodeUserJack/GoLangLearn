package main

import (
	"fmt"
	"errors"
	"os"
	"time"
)

// custom error
type cerror struct {
	name string
	time time.Time
	severity int
}

func (c cerror) String() string {
	s := c.name
	s = s + " "
	s += c.time.Format("2006-01-02 15:04:05.000000")
	s = s + " "
	s += string(c.severity)
	return s
}

func (c cerror) Error() string {
	res := cerror {
		name: "Custom Error",
		time: time.Now(),
		severity: 5,
	}
	return res.String()
}

func main() {
	i, err := fmt.Println("Error check!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("bytes Printed: ", i)

	// fmt scan
	var name string
	var age int

	n, err := fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("What is your age?")
	fmt.Scan(&age)
	fmt.Println(name, age, n, err)

	// recover
	f()
	fmt.Println("f returned normally")

	// error info / custom error
	val, err := errorNew(55)
	if err != nil {
		fmt.Println(val, err)
	}

	// tangential topic of os.Args for commandline args, will have .exe file name
	args := os.Args
	fmt.Println(args)

	// Errorf()
	vv, erro := errorF(66)
	if erro != nil {
		fmt.Println(vv, erro)
	}

	// custom error
	val, errr := ccerror(0)
	if errr != nil {
		fmt.Println(val, errr)
	}
}

// recover
func f() {
	defer func(){
		if r:=recover(); r != nil {
			fmt.Println("Recovered in f")
		}
	}()
	g(3)
	fmt.Println("not printed")
}

func g(i int) {
	if i >= 3 {
		panic("panic started")
	}
	defer func(){
		fmt.Println("g defer panicking")
	}()
}

// error.New()
func errorNew(i int) (int, error) {
	if i > 50 {
		return 50, errors.New("i bigger than 50")
	}
	return i / 2, nil
}

// errorF()
func errorF(x int) (int, error) {
	if x < 100 {
		return x, fmt.Errorf("%d is less than 100", x)
	}
	return x - 100, nil
}

// custom error
func ccerror(y int) (int, error) {
	if y == 0 {
		return y, cerror {"y is 0", time.Now(), 3}
	}
	return y * 2, nil
}