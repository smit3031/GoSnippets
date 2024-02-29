package basics

import "fmt"

func RunOperaters() {
	fmt.Println("Welcome to operaters!")

	// Airthmetic Ops
	p := 34
	q := 20

	// Addition
	result1 := p + q
	fmt.Printf("Result of p + q = %d", result1)

	// Subtraction
	result2 := p - q
	fmt.Printf("\nResult of p - q = %d", result2)

	// Multiplication
	result3 := p * q
	fmt.Printf("\nResult of p * q = %d", result3)

	// Division
	result4 := p / q
	fmt.Printf("\nResult of p / q = %d", result4)

	// Modulus
	result5 := p % q
	fmt.Printf("\nResult of p %% q = %d", result5)

	//Relational Ops
	a := 34
	b := 20

	// ‘=='(Equal To)
	res1 := a == b
	fmt.Println(res1)

	// ‘!='(Not Equal To)
	res2 := a != b
	fmt.Println(res2)

	// ‘<‘(Less Than)
	res3 := a < b
	fmt.Println(res3)

	// ‘>'(Greater Than)
	res4 := a > b
	fmt.Println(res4)

	// ‘>='(Greater Than Equal To)
	res5 := a >= b
	fmt.Println(res5)

	// ‘<='(Less Than Equal To)
	res6 := a <= b
	fmt.Println(res6)

	//Logical ops
	var x int = 23
	var y int = 60

	if x != y && x <= y {
		fmt.Println("True")
	}

	if x != y || x <= y {
		fmt.Println("True")
	}

	if !(x == y) {
		fmt.Println("True")
	}

	//Bitwise Ops
	v1 := 34
	v2 := 20

	// & (bitwise AND)
	r1 := v1 & v2
	fmt.Printf("Result of v1 & v2 = %d", r1)

	// | (bitwise OR)
	r2 := v1 | v2
	fmt.Printf("\nResult of v1 | v2 = %d", r2)

	// ^ (bitwise XOR)
	r3 := v1 ^ v2
	fmt.Printf("\nResult of v1 ^ v2 = %d", r3)

	// << (left shift)
	r4 := v1 << 1
	fmt.Printf("\nResult of v1 << 1 = %d", r4)

	// >> (right shift)
	r5 := v1 >> 1
	fmt.Printf("\nResult of v1 >> 1 = %d", r5)

	// &^ (AND NOT)
	r6 := v1 &^ v2
	fmt.Printf("\nResult of v1 &^ v2 = %d", r6)

	//Assignment Ops
	var u int = 45
	var v int = 50

	// “=”(Simple Assignment)
	u = v
	fmt.Println(u)

	// “+=”(Add Assignment)
	u += v
	fmt.Println(u)

	//“-=”(Subtract Assignment)
	u -= v
	fmt.Println(u)

	// “*=”(Multiply Assignment)
	u *= v
	fmt.Println(u)

	// “/=”(Division Assignment)
	u /= v
	fmt.Println(u)

	// “%=”(Modulus Assignment)
	u %= v
	fmt.Println(u)

	//Misc Ops
	// Using address of operator(&) and
	// pointer indirection(*) operator
	var1 := 3
	var2 := &var1
	fmt.Println(*var2)
	*var2 = 7
	fmt.Println(var1)
}
