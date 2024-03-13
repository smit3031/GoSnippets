package arrays

import "fmt"

func RunArrays() {
	fmt.Println("Hey From Arrays[]!!!")

	/*
		An array is a fixed-length sequence that is used to store homogeneous elements in the memory.

		Var array_name[length]Type

		array_name:= [length]Type{item1, item2, item3,...itemN}
	*/

	var ar [2]string
	ar[0] = "xy"
	ar[1] = "wz"

	fmt.Println(ar)

	arr2 := [4]int{2, -1, 3, 4}
	fmt.Println(arr2)

	//Multi-Dimensional Array

	/*
		Multi-Dimensional arrays are the arrays of arrays of the same type.
		Array_name[Length1][Length2]..[LengthN]Type

		In a multi-dimension array, if a cell is not initialized with some value
		by the user, then it will initialize with zero by the compiler automatically.
		There is no uninitialized concept in the Golang.
	*/

	arr3 := [3][3]string{{"RCB", "CSK", "MI"}, {"FCB", "RM", "LFC"},
		{"IPL", "UCL", "EPL"}}

	fmt.Println("Elements of Array 3")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr3[x][y])
		}
	}

	arr4 := [3][2][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
		{
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		{
			{17, 18, 19, 20},
			{21, 22, 23, 24},
		},
	}
	fmt.Println(arr4, "   Length : ", len(arr4))

	//Defaults
	var myarr [3]int
	fmt.Println(myarr)

	var myarr2 [2]string
	fmt.Println(myarr2)
	fmt.Println(len(myarr2))

	// In an array, if ellipsis ‘‘…’’ become visible at the place of length,
	// then the length of the array is determined by the initialized elements.
	myarray := [...]int{2, 4, 6, 8}

	fmt.Println("Elements of the array: ", myarray, ", Length :", len(myarray))

	//Iterate with for loop
	for i := 0; i < len(myarray); i++ {
		fmt.Println(myarray[i] / 2)
	}

	arr := [5]int{1, 2, 3, 4, 5}

	// 1. Using a Traditional For Loop
	fmt.Println("1. Traditional For Loop:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 2. Using Range in For Loop
	fmt.Println("\n2. Using Range in For Loop:")
	for i, value := range arr {
		fmt.Println(i, value)
	}

	// 3. Using Blank Identifier for Unused Index/Value
	fmt.Println("\n3. Using Blank Identifier:")
	for _, value := range arr {
		fmt.Println(value)
	}

	// 4. Using Only Index with Range
	fmt.Println("\n4. Using Only Index with Range:")
	for i := range arr {
		fmt.Println(i)
	}

	// 5. Using Only Value with Range
	fmt.Println("\n5. Using Only Value with Range:")
	for _, value := range arr {
		fmt.Println(value)
	}

	// 6. Using While-Style For Loop
	fmt.Println("\n6. While-Style For Loop:")
	i := 0
	for i < len(arr) {
		fmt.Println(arr[i])
		i++
	}

	/*
		In Go language, an array is of value type not of reference type.
		So when the array is assigned to a new variable,
		then the changes made in the new variable do not affect the original array.
	*/

	my_array := [...]int{100, 200, 300, 400, 500}
	fmt.Println("Original array(Before):", my_array)

	new_array := my_array

	fmt.Println("New array(before):", new_array)

	new_array[0] = 500

	fmt.Println("New array(After):", new_array)

	fmt.Println("Original array(After):", my_array)

	/*
		In an array, if the element type of the array is comparable,
		then the array type is also comparable. So we can
		directly compare two arrays using == operator.
	*/
	ar1 := [3]int{9, 7, 6}
	ar2 := [...]int{9, 7, 6}
	ar3 := [3]int{9, 9, 9}

	// Comparing arrays using == operator
	fmt.Println(ar1 == ar2)
	fmt.Println(ar2 == ar3)
	fmt.Println(ar1 == ar3)

}
