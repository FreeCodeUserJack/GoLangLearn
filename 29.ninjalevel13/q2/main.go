package main

import(
	"fmt"
	"github.com/FreeCodeUserJack/GoLangLearn/29.ninjalevel13/q2/quote"
	"github.com/FreeCodeUserJack/GoLangLearn/29.ninjalevel13/q2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Printf("key: %v val: %v, ", k, v)
	}
	fmt.Println()
}