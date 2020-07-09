package acdc

import (
	"fmt"
	"testing"
)

func BenchmarkSum(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}...)
	}
}

func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{
			data:   []int{1, 2, 3, 4, 5},
			answer: 15,
		},
		test{
			data:   []int{9, 10, 11},
			answer: 30,
		},
		test{[]int{5, 4, 3, 2, 1}, 15},
		test{[]int{-1, 0, 1}, 0},
	}

	res := Sum(1, 2, 3, 4, 5)
	if res != 15 {
		t.Error("Expected 15 but got: ", res)
	}

	for _, val := range tests {
		ans := Sum(val.data...)
		if ans != val.answer {
			t.Errorf("Expected %d but got %d", val.answer, ans)
		}
	}
}
