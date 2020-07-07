package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
		for j := 0; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	j := 1
	for j < 10 {
		fmt.Println("for loops are cool!")
		j += 5
	}

	var k int = 33
	for k <= 122 {
		fmt.Printf("%d\t%+q\t%#U\n", k, k, k)
		k++
	}
	fmt.Printf("%v %+q %#U %U", 300, 300, 300, 300)
}