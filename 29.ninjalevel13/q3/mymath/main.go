package mymath

import "sort"

// CenteredAvg ...computes average of a list of int
// after removing min and max
func CenteredAvg(data []int) float64 {
	sort.Ints(data)
	use := data[1:len(data)-1]

	sum := 0
	for _, v := range(use) {
		sum += v
	}

	res := float64(sum) / float64(len(use))
	return res
}