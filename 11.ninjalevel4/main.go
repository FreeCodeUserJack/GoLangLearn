package main

import (
	"fmt"
)

func main() {
	// #1 & 2
	cla := [5]int{0, 2, 3, 4, 5}
	cls := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	cla[0] = 1
	for i, val := range cla {
		fmt.Println(i, val)
	}
	for _, val := range cla {
		fmt.Printf("%T ", val)
	}
	fmt.Printf("\n%T\n", cla)
	fmt.Printf("%T\n", cls)

	// #3
	fmt.Println(cls[5:])
	fmt.Println(cls[:5])
	
	// #4
	cls = append(cls, 52)
	fmt.Println(cls)
	cls = append(cls, 53, 54, 55)
	x := []int{56, 57, 58}
	cls = append(cls, x...)
	fmt.Println(cls)

	// #5
	cls = append(cls[:10], cls[12:]...)
	fmt.Println(cls)

	// #6
	states := make([]string, 5, 51) //cap will be 51 as well
	fmt.Println(states, len(states), cap(states))
	data := []string{"Arizona"} // paste the data here
	states = append(states, data...)
	fmt.Println(states)
	for in:=0; in < len(states); in++ {
		fmt.Println(in, states[in])
	}

	// #7
	slice := make([][]string, 2)
	for k:=0; k<len(slice); k++ {
		slice[k] = make([]string, 0, 4)
	}
	slice[0] = []string{"a", "b", "c", "d"}
	slice[1] = []string{"1", "2", "3", "4"}
	fmt.Println(slice)
	for _, val := range slice {
		for _, valu := range val {
			fmt.Println(valu)
		}
	}

	// #8
	mm := map[string] []string{}
	mm["Jack"] = []string{"a", "b", "c"}
	mm["Jill"] = []string{"1", "2", "3"}
	fmt.Println(mm)

	for k, v := range mm {
		for pos, val := range v {
			fmt.Println(k, pos, val)
		}
	}

	// #9
	mm["feawe"] = []string{"f", "few", "tt"}
	for k, v := range mm {
		fmt.Println(k, v)
	}

	// #10
	delete(mm, "feawe")
	fmt.Println(mm)
}