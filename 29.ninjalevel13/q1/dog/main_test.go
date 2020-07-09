package dog

import (
	"fmt"
	"testing"
)

// go test
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out
// go tool -bench .

func BenchmarkYears(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Years(11111)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i:=0; i<b.N; i++ {
		YearsTwo(11111)
	}
}

// table testing
type test struct {
	input int
	output int
}

func TestYears(t *testing.T) {
	tests := []test{
		test{
			input: 5,
			output: 35,
		},
		test{1, 7},
		test{2, 14},
		test{10, 70},
	}

	for _, val := range tests {
		res := Years(val.input)
		if res != val.output {
			t.Errorf("Expected %v, got %d", val.output, res)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	tests := []test{
		test{
			input: 5,
			output: 35,
		},
		test{1, 7},
		test{2, 14},
		test{10, 70},
	}

	for _, val := range tests {
		res := YearsTwo(val.input)
		if res != val.output {
			t.Errorf("Expected %v, got %d", val.output, res)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output: 49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}