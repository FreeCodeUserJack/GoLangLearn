package main

import (
	"fmt"
	"encoding/json"
	"os"
	"io"
	"sort"
	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string `json:"first"`
	Last string `json:"last"`
	Age int `json:"age,omitempty"`
}

// to sort by Age
type byAge []person

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

// to sort by Name
type byName []person

func (p byName) Len() int {
	return len(p)
}

func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byName) Less(i, j int) bool {
	return p[i].First < p[j].First
}

func main() {
	p1 := person {
		First: "James",
		Last: "Sheen",
		Age: 55,
	}

	p2 := person {
		First: "Jane",
		Last: "Fonda",
		Age: 50,
	}

	persons := []person{p1, p2}
	fmt.Println(persons)

	b, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// unmarshal the output of the println(b)
	jp := `[{"First":"James","Last":"Sheen","Age":55},{"First":"Jane","Last":"Fonda","Age":50}]`
	bs := []byte(jp)

	fmt.Printf("%T %T\n", jp, bs)

	// resdata := []person{}
	var resdata []person
	err1 := json.Unmarshal(bs, &resdata)
	if err != nil {
		fmt.Println(err1)
	}
	fmt.Println(resdata)

	// Writer Interface
	fmt.Fprintln(os.Stdout, "Hello from Fprintln")
	io.WriteString(os.Stdout, "Hello from io.WriteString")

	// sorting a slice
	x := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	y := []string{"James", "Q", "M", "MoneyPenny", "Dr. No"}

	sort.Ints(x)
	sort.Strings(y)
	fmt.Println(x, y)

	// custom sort
	pp1 := person{"John", "Doe", 22}
	pp2 := person{"Jane", "Doe", 30}
	pp3 := person{"Jessie", "Jane", 40}
	pp4 := person{"Jack", "Atlas", 18}

	people := []person{pp1, pp2, pp3, pp4}
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println(people)

	// bcrypt
	pass := `password123`
	cipher, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cipher)

	erro := bcrypt.CompareHashAndPassword(cipher, []byte(pass))
	if erro != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Passwords are the same")
	}

}