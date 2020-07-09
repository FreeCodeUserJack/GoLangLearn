package mymath

import (
	"fmt"
	"testing"
)
func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		ans float64
	}

	tests := []test{
		test {
			data: []int{1, 2, 3, 4, 5},
			ans: 3,
		},
		test {[]int{1, 2, 3}, 2},
		test {[]int{9999, 55, 4, 57}, 56},
	}

	for _, tst := range tests {
		res := CenteredAvg(tst.data)
		if tst.ans != res {
			t.Errorf("Expected %f, got %f", tst.ans, res)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output: 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i:=0; i<b.N; i++ {
		CenteredAvg([]int{12, 55, 65, 31, 5454, 48, 909, 444, 3333, 234})
	}
}