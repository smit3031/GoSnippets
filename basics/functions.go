package basics

import (
	"fmt"
	"math"
	"strings"
)

func RunFunctions() {

	/*
				func function_name(Parameter-list)(Return_type){
		    		// function body.....
				}
	*/

	fmt.Println(math.Abs(-2), area(3, 2))

	var a = 0
	var b = 1
	//Call By Value
	fmt.Println("Before swapping a and b : ", a, b)
	swap(a, b)
	fmt.Println("After swapping a and b :  ", a, b)

	//Call By Reference
	fmt.Println("Before swapping a and b : ", a, b)
	swapByRef(&a, &b)
	fmt.Println("After swapping a and b :  ", a, b)

	//Variadic Functions in Go

	/*
		The function that is called with the varying number
		of arguments is known as variadic function.

		In the declaration of the variadic function,
		the type of the last parameter is preceded by an ellipsis, i.e, (…).
		It indicates that the function can be called
		at any number of parameters of this type.

		Syntex:
		function function_name(para1, para2...type)type {// code...}

		The variadic functions are generally used for functions that perform string formatting.
	*/
	fmt.Println(joinstr())
	fmt.Println(joinstr("Smit", "Chaudhari"))
	fmt.Println(joinstr("Chandler", "Bing", "Or", "Bong"))

	// Anonymous functions in go

	/*
		Go language provides a special feature known as an anonymous function. An anonymous function is a function which doesn’t contain any name. It is useful when you want to create an inline function. In Go language, an anonymous function can form a closure. An anonymous function is also known as function literal. Syntax:
		func(parameter_list)(return_type){
			code..

			Use return statement if return_type are given
			if return_type is not given, then do not
			use return statement
			return
		}()
	*/

	func() {
		fmt.Println("Welcome! to ghost function!! wink!!")
	}()

	/*
		In Go language, you are allowed to assign an anonymous function to a variable.
		When you assign a function to a variable, then the type of the variable is of
		function type and you can call that variable like a function call
	*/
	value := func() {
		fmt.Println("Welcome! to value variable as ghost function!! wink!!")
	}
	value()

	// Passing arguments in anonymous function
	func(ele string) {
		fmt.Println(ele)
	}("Passed ele string as parameter in ghost function!! wink!!")

	//You can also pass an anonymous function as an argument into other function.
	val := func(p, q string) string {
		return p + q + "Bong"
	}
	ghostFunc(val)

	// You can also return an anonymous function from another function.
	val2 := returnGhostFunc()
	fmt.Println(val2("Chandler ", "Bing "))

	//Main Function
	/*
		It is entry point of program and every executable program should
		have only one main package and function.
		It does not take any argument nor return anything.
	*/

	//Init Function
	/*
		init() function is just like the main function,
		does not take any argument nor return anything.
		This function is present in every package and this function is called when
		the package is initialized.

		This function is declared implicitly, so you cannot
		reference it from anywhere and you are allowed to create multiple
		init() function in the same program and they execute in the order they are created.

		And allowed to put statements if the init() function, but always remember
		to init() function is executed before the main() function call,
		so it does not depend to main() function. The main purpose of the init() function
		is to initialize the global variables that cannot be initialized in the global context.
	*/

	//Blank Identifier
	/*
		_ is used when a function returns multiple values and we don't want to use certain value.
		Otherwise in go if we don't use defined variable compiler throws error.
	*/
	_, a, str := returnMultiVals()
	fmt.Println("Second value (int):", a)
	fmt.Println("Third value (string):", str)
}

func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello! init() function")
}

func area(length, width int) int {
	Ar := length * width
	return Ar
}

func swap(a, b int) {

	var o int
	o = a
	a = b
	b = o

}

func swapByRef(a, b *int) int {
	var o int
	o = *a
	*a = *b
	*b = o

	return o
}

func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func ghostFunc(i func(p, q string) string) {
	fmt.Println(i("Chandler", "Bing"))
}

// Returning anonymous function
func returnGhostFunc() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "Bong"
	}
	return myf
}

func returnMultiVals() (int, int, string) {
	return 2, 3, "Let's ignore 2 totally"
}
