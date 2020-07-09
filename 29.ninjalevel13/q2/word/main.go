package word

import "strings"

// UseCount ...will return map of count of occurrences of each word
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, val := range(xs) {
		m[val]++
	}
	return m
}

// Count ...will return number of words in entire string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}