package main

import "fmt"

// #1
type person struct {
	first string
	last string
	iceCream []string
}

func main() {
	// #1
	p1 := person {
		first: "James",
		last: "Johnson",
		iceCream: []string{"choco", "vanil"},
	}
	fmt.Println(p1.first, p1.last)
	for _, val := range p1.iceCream {
		fmt.Printf("\t%s\n", val)
	}

	// #2
	mp := map[string] person{}
	mp[p1.last] = p1
	for _, v := range mp {
		fmt.Println(v.first, v.last)
		for _, val := range v.iceCream {
			fmt.Printf("\t%s\n", val)
		}
	}

	// #3
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck {
		vehicle: vehicle {
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}
	s1 := sedan {
		vehicle: vehicle {
			doors: 2,
			color: "green",
		},
		luxury: true,
	}
	fmt.Println(t1, s1)
	fmt.Println(t1.fourWheel, s1.luxury)

	// #4
	anon := struct {
		blah string
	} {
		blah: "blah blah",
	}
	fmt.Println(anon.blah)
}