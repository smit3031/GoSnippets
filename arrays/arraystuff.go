package arrays

import "fmt"

func RunArraysStuff() {
	fmt.Println("Heya! arrays here!")

	/*
		Creating a copy of an array by value
		arr := arr1
		Creating a copy of an array by reference
		arr := &arr1
	*/

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr3 := &arr1

	fmt.Println(arr1, " , ", arr2, " , ", arr3)

	arr3[2] = -3
	arr2[2] = -9
	fmt.Println(arr1, " , ", arr2, " , ", arr3)

	/*
		Various methods to copy arrays.
	*/
	originalArray := []int{1, 2, 3, 4, 5}
	copyArray := make([]int, len(originalArray))

	for i, value := range originalArray {
		copyArray[i] = value
	}

	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copy Array: ", copyArray)

	copyArray2 := make([]int, len(originalArray))
	copy(copyArray2, originalArray)
	fmt.Println("Copy Array2: ", copyArray2)

	copyArray3 := make([]int, 0, len(originalArray))
	copyArray3 = append(copyArray3, originalArray...)
	fmt.Println("Copy Array3: ", copyArray3)

	//Pass array in function
	/*
		// For sized array
		func function_name(variable_name [size]type){
		// Code
		}

		In Go, arrays are passed to functions as values, not as references.
		This means that changes made to the array inside the function will not affect the original array.
		To pass an array to a function in Go, you simply need to pass the array as an argument to the function.
	*/
	ans := myFun([5]int{1, 2, 3, 4, 5})
	ans++
}

func myFun(arr [5]int) int {
	var ans, res int
	res = 2
	for i := 0; i < 5; i++ {
		ans += arr[i]
	}
	return ans / res
}
