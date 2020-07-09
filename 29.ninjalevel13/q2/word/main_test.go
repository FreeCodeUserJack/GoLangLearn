package word

import(
	"fmt"
	"testing"
	"github.com/FreeCodeUserJack/GoLangLearn/29.ninjalevel13/q2/quote"
)

// Count
func BenchmarkCount(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("This is a sentence"))
	// Output: 4
}

func TestCount(t *testing.T) {
	type test struct {
		s string
		ans int
	}

	tests := []test{
		test {
			s: "This is fine",
			ans: 3,
		},
		test {"This is", 2},
		test {"This", 1},
		test {"", 0},
	}

	for _, val := range tests {
		res := Count(val.s)
		if res != val.ans {
			t.Errorf("Expected %d, got %d", val.ans, res)
		}
	}
}

// UseCount
func BenchmarkUseCount(b *testing.B) {
	for i:=0; i<b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("This is input"))
	// Output:
	// map[This:1 input:1 is:1]
}

func TestUseCount(t *testing.T) {
	type test struct {
		s string
		ans map[string]int
	}

	tests := []test{
		test {
			s: "This is fine",
			ans: map[string]int{"This": 1, "is": 1, "fine":1,},
		},
		test {"This", map[string]int{"This":1,}},
		test {"", map[string]int{}},
	}

	for _, val := range tests {
		res := UseCount(val.s)
		for k, v := range res {
			if val.ans[k] != v {
				t.Errorf("Expected %v, got %v", val.ans[k], v)
			}
		}
	}
}