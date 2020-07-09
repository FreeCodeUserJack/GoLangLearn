package dog

import (
	"fmt"
)

func main() {
	fmt.Println("nil")
}

// Years ...this function will input int of human years and then
// return the equivalent in dog years.
func Years(h int) int {
	return int(h / 7)
}