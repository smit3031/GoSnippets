package basics

import "fmt"

func AllDataTypes() {

	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	fmt.Println("here are all data types")

	a := 20.45
	b := 34.89

	// Subtraction of two
	// floating-point number
	c := b - a

	// Display the result
	fmt.Printf("Result is: %f", c)

	// Display the type of c variable
	fmt.Printf("\nThe type of c is : %T", c)

	var x float32 = 5.00
	var y float32 = 2.25
	//Addition
	fmt.Printf("addition :  %g + %g = %g\n ", x, y, x+y)
	//Subtraction
	fmt.Printf("subtraction : %g - %g = %g\n", x, y, x-y)
	//Multiplication
	fmt.Printf("multiplication : %g * %g = %g\n", x, y, x*y)
	//Division
	fmt.Printf("division : %g / %g = %g\n", x, y, x/y)

	var p complex128 = complex(6, 2)
	var q complex64 = complex(9, 2)
	fmt.Println(p)
	fmt.Println(q)

	// Display the type
	fmt.Printf("The type of a is %T and "+
		"the type of b is %T", a, b)

	comp1 := complex(10, 11)
	// complex number init syntax
	comp2 := 13 + 33i
	fmt.Println("Complex number 1 is :", comp1)
	fmt.Println("Complex number 1 is :", comp2)
	// get real part
	realNum := real(comp1)
	fmt.Println("Real part of complex number 1:", realNum)
	// get imaginary part
	imaginary := imag(comp2)
	fmt.Println("Imaginary part of complex number 2:", imaginary)

	// str variable which stores strings
	str := "GeeksforGeeks"

	// Display the length of the string
	fmt.Printf("Length of the string is:%d",
		len(str))

	// Display the string
	fmt.Printf("\nString is: %s", str)

	// Display the type of str variable
	fmt.Printf("\nType of str is: %T", str)
}
