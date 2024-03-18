package arrays

import "fmt"

func RunSlices() {
	fmt.Println("Hey! Go Spread butter on my Slice!")

	/*
		Slices in Go are a flexible and efficient way to represent arrays, and
		they are often used in place of arrays because of their dynamic size and added features.
		A slice is a reference to a portion of an array.

		Slice is a variable-length sequence that stores elements of a similar type,
		you are not allowed to store different type of elements in the same slice.
		It is just like an array having an index value and length,
		but the size of the slice is resized they are not in fixed-size just like an array.
	*/

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	fmt.Println("Array: ", array)
	fmt.Println("Slice: ", slice)

	slice2 := append(slice, 2, -6, 1)
	fmt.Println("New Slice", slice2)

	// pprof -> try
}
