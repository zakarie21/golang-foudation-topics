package main

import "fmt"

func main() {
	/*
		- using two separate for loops, generate two different slices
	- show how you can merge these slices together, one way, and the other way
	- show a situation where editing a slice will also edit another slice, and show another situation where it won't
	- (bonus) create a two-dimensional slice (slice within a slice!)


	*/
	slice1 := []int{}

	slice2 := []int{}

	editSlice1 := []int{3, 6, 9, 12}

	for i := 0; i <= 30; i += 5 {
		slice1 = append(slice1, i)
	}

	for i := 0; i <= 10; i++ {
		slice2 = append(slice2, i)
	}

	fmt.Println(slice1)
	fmt.Println("")
	fmt.Println(slice2)
	fmt.Println("")

	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	fmt.Println("")

	for i := 0; i <= 20; i += 2 {
		editSlice1 = append(editSlice1, i)
	}

	editSlice2 := editSlice1

	editSlice1[0] = 70

	fmt.Println(editSlice1)
	fmt.Println("")
	fmt.Println(editSlice2)
	fmt.Println("")

	editSlice1 = append(editSlice1, 500)

	fmt.Println(editSlice1)
	fmt.Println("")
	fmt.Println(editSlice2)
	fmt.Println("")

	runtimes := [][]float32{
		{10.4, 9.7, 11.0},
		{9.4, 10.3, 11.6},
		{10.1, 9.9, 12.3},
	}

	fmt.Println(runtimes)
}
