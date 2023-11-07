package main

import "fmt"

func main() {

	var a [4][5]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = i + j
		}
	}
	fmt.Println("2D Array:", a)

	//  x := []int{10, 20, 30, 40, 50}
	//  for index := 0; index < len(x); index++ {
	//   fmt.Printf("x[%d] = %d\n", index, x[index])
	//  }

	x := []int{10, 20, 30, 40, 50}
	for index, value := range x {
		fmt.Printf("x[%d] = %d\n", index, value)
	}

	// Creating an array of size 8 and slice this array till 5
	// and return the reference of the slice using make function
	// var my_slice = make([]int, 5, 8)
	// fmt.Printf("Slice = %v", my_slice)
	// fmt.Printf("\nLength = %d", len(my_slice))
	// fmt.Printf("\nCapacity = %d", cap(my_slice))

	// // reference type are slices:
	// a := [7]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven"}

	// slice1 := a[1:]
	// slice2 := a[3:]

	// fmt.Println("***** Before Modifications *****")
	// fmt.Println("Array =", a)
	// fmt.Println("slice1 =", slice1)
	// fmt.Println("slice2 =", slice2)

	// slice1[0] = "TWO"
	// slice1[2] = "FOUR"

	// slice2[1] = "FIVE"

	// fmt.Println("\n***** After Modifications *****")
	// fmt.Println("Array =", a)
	// fmt.Println("slice1 =", slice1)
	// fmt.Println("slice2 =", slice2)
}
