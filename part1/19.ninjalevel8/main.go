package main

import (
	"fmt"
	"encoding/json"
	"os"
	"sort"
)

// #1
type user struct {
	First string `json:"first"`
	Age int `json:"age,omitempty"`
}

// #5
type byAge []user
func(u byAge) Len() int {
	return len(u)
}
func(u byAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func(u byAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

type byName []user
func(u byName) Len() int {
	return len(u)
}
func(u byName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func(u byName) Less(i, j int) bool {
	return u[i].First < u[j].First
}

// #4
type byInt []int

func (b byInt) Len() int {
	return len(b)
}
func (b byInt) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byInt) Less(i, j int) bool {
	return b[i] < b[j]
}

func main() {
	// #1
	u1 := user{
		First: "James",
		Age: 32,
	}
	u2 := user{"MoneyPenny", 25,}
	u3 := user{
		First: "Ajax",
		Age: 28,
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)

	res, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

	// #2
	var newusers []user
	// newusers := []user{}
	data := `[{"first":"James","age":32},{"first":"MoneyPenny","age":25},{"first":"Ajax","age":28}]`
	json.Unmarshal([]byte(data), &newusers)
	fmt.Println(newusers)
	// if given JSON format, can use online tool to conert JSON to Go Struct code

	// #3
		// encode to JSON []user and send to Stdout -> need json.NewEncoder(os.Stdout).encode(v interface{})
	eror := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(eror)
	}

	// #4
	xi := []int{4, 6, 23, 64, 13, 76, 23, 98, 67, 34, 876, 34, 1}
	xs := []string{"b", "awe", "nm", "fc", "ccea", "awefwefaw", "efefe"}

	fmt.Println(xi, xs)
	// sort.Ints(xi)
	sort.Sort(byInt(xi))
	sort.Strings(xs)
	fmt.Println(xi, xs)

	// #5
		// sort users by age and first, see above
	sort.Sort(byAge(users))
	fmt.Println(users)
	sort.Sort(byName(users))
	fmt.Println(users)
}