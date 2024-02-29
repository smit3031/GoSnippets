package basics

import "fmt"

// Geeks, geeks, _geeks23  // valid variable

/*
	var variable_name type = expression

	ex.   var x int = 2
	OR    var x = 2
	OR    x := 2

	If value is not assigned default(0 in this case) will be assigned
*/

func RunVariables() {
	fmt.Println("vars")
	var myvariable1 int
	var myvariable2 string
	var myvariable3 float64

	// Display the zero-value of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n",
		myvariable2)

	fmt.Printf("The value of myvariable3 is : %f",
		myvariable3)

	var x, y, z int = 1, 2, 4
	fmt.Println("   ", x, y, z)

	var a, b, c = "Smit", 3.44, 4
	fmt.Println(a, b, c)

	myvar1, myvar2, myvar3 := 800, 34, 56
	fmt.Println(myvar1, myvar2, myvar3)
}
