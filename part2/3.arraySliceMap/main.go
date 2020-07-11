package main

import (
	"fmt"
)

func main() {
	// arr, type is [5]int, arrays not used in Go, use slices
	var a1 [5]int
	fmt.Println(a1)
	a1[2] = 100
	fmt.Println(a1)

	// slice
	var a2 []int
	fmt.Println(a2, a2 == nil)
	a3 := []int{1, 2, 3}
	a2 = append(a2, a3...)
	fmt.Println(a2, len(a2), "cap: ", cap(a2))

	// delete from slice
	a2 = append(a2[0:1], a2[2:]...)
	fmt.Println(a2)

	// multi dimensional slice
	var mds [][]int = make([][]int, 5)
	for xx := 0; xx < len(mds); xx++ {
		mds[xx] = make([]int, 5)
	}
	fmt.Println(mds)
	newRow := []int{1, 2, 3, 4, 5}
	mds = append(mds, newRow)
	for _, row := range mds {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

	// map
	mp := map[string]int{
		"Jack": 1,
		"Jill": 2,
	}
	// never do this b/c it's nil and you can't append to it
	var mppN map[string]int
	fmt.Println(mppN, mppN == nil)

	mpp := map[string]int{}
	fmt.Println(mpp, mpp == nil)
	fmt.Printf("%v, %T\n", mp, mp)
	var maap map[int]string = make(map[int]string)
	fmt.Println(maap, maap == nil)
	maap[1] = "Jack"
	maap[2] = "Jill"
	for k, v := range maap {
		fmt.Println(k, v)
	}
	name, ok := maap[3]
	if !ok {
		fmt.Println("index 3 not found: ", name)
	}
	maap[3] = "James"
	delete(maap, 1)
	fmt.Println(maap)
}
