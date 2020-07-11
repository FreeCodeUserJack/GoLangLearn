package main

import "fmt"

func main() {

	// multi const declaration and assignment
	const (
		A = iota * 10
		B = iota * 10
		C int = iota * 10
		D = 1 << 5 - 2
	)
	fmt.Println(A, B, C, D)

	// pointer
	var a int = 5
	var b *int = & a
	*b = 6
	fmt.Println(b, a)

	// loop
	for i:=0; i<10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	j := 0
	for j < 11 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()
	k := -1
	for {
		k++
		if k % 2 == 0 {
			continue
		}
		if k > 20 {
			break
		}
		fmt.Print(k, " ")
	}
	fmt.Println()
	for a:=0; a<10; a++ {
		for b:=0; b<10; b++ {
			fmt.Printf("%d:%d ", a, b)
		}
		fmt.Println()
	}

	// conversion
	str := "Hello 世界" // note how each of chinese chars are using 3 bytes
	fmt.Println([]byte(str))
	fmt.Println(string([]byte(str)))
	fmt.Println(string([]byte{228, 184, 150, 231, 149, 140}))
	fmt.Println([]rune(str))

	// switch
	swi := "en"
	switch swi {
	case "en":
		fmt.Println("english!")
		fallthrough
	case "es":
		fmt.Println("spanish!")
	case "ge", "de":
		fmt.Println("german!")
	default:
		fmt.Println("not en or es!")
	}

	// switch on type
	var bb interface{} = "string"
	switch t := bb.(type) {
	case int, int16, int32, int64:
		fmt.Println("int!", t)
	case string:
		fmt.Println("string!")
	case float64, float32:
		fmt.Println("float!")
	default:
		fmt.Println("NIL!")
	}
}