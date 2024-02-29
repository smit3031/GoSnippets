package basics

import (
	"fmt"
	"math"
)

// you can use bracktes to intialise multiple constants
const (
	p = "GeeksforGeeks"
	q = true
	r = 3.14
)

func RunConsts() {

	const PI = 3.14
	const GFG = "GeeksforGeeks"
	fmt.Println("Hello", GFG)

	fmt.Println("Happy", PI, "Day")

	const Correct = true
	fmt.Println("Go rules?", Correct)

	const n = 5

	const d = 3e10 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	fmt.Println(p, q, r)
}
