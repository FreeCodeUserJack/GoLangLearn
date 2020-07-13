package main

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
	"sort"
)

// House ...will represent a house
type House struct {
	Cost float64 `json:"cost,omitempty"`
	Size float64 `json:"size"`
	Name string `json:"name"`
}

// Person ...will represent a person
type Person struct {
	House House // embedding (composition), promotion will make this pseudo-inheritance
	Name string `json:"name"`
	Password string `json:"-"`
}

func (h House) String() string {
	return fmt.Sprintf("Cost: %f, Size: %f, Name: %s", h.Cost, h.Size, h.Name)
}

func (p Person) speak() {
	fmt.Println("Hello there from Person")
	fmt.Println("this is my house: ", p.House)
}

type vocal interface {
	makeNoise()
}

func (p Person) makeNoise() {
	fmt.Println("This is human noise")
}

func info(v vocal) {
	fmt.Print("let's make some noise! ")
	v.makeNoise()
}

type byName []Person

func (b byName) Len() int {
	return len(b)
}
func (b byName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}
func (b byName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	p1 := Person {
		House: House {
			Size: 2222,
			Name: "home1",
		},
		Name: "John Doe",
		Password: "pass123",
	}
	ptr1 := &p1
	fmt.Printf("%T %v\n", p1, p1)
	fmt.Printf("%T %v %p\n", ptr1, ptr1, ptr1)

	// marshal
	payload, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(payload))

	// unmarshal
	str := `{"House":{"size":2222,"name":"home1"},"name":"John Doe","password":"pass123"}`
	var p2 Person
	err2 := json.Unmarshal([]byte(str), &p2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%+v\n", p2)

	// encode
	json.NewEncoder(os.Stdout).Encode(p1)

	// decode
	var p3 Person
	read := strings.NewReader(str)
	json.NewDecoder(read).Decode(&p3)
	fmt.Printf("%T %v %p\n", p3, p3, &p3)

	// interface methods
	fmt.Printf("\n############################ Inteface Below\n")
	p1.speak()
	p1.makeNoise()
	info(p1)

	// sort struct
	p4 := Person {
		House: House {
			Cost: 250000.0,
			Size: 2222,
			Name: "home1",
		},
		Name: "Jand Doe",
		Password: "pass123",
	}
	ppl := []Person{p1, p4}
	fmt.Println(ppl)
	sort.Sort(byName(ppl))
	fmt.Println(ppl)
	fmt.Println(p4.House)

	// reverse ints array sort
	ints := []int{4, 3, 6, 3, 9, 1, 99, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)
}