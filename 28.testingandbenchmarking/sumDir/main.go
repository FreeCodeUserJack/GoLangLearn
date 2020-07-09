package acdc

import (
	// "fmt"
)

func main() {
	
	// fmt.Println("sum of 1, 2, 3, 4, 5: ", Sum(1, 2, 3, 4, 5))
	// fmt.Println("sum of 1, 2, 3, 4, 5, 6 : ", Sum(1, 2, 3, 4, 5, 6))

}

// Sum ...will take variadic ints and returns sum
func Sum(x ...int) int {
	sum := 0
	for _, val := range x {
		sum += val
	}
	return sum
}
