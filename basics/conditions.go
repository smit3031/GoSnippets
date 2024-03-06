package basics

import "fmt"

func RunConditions() {
	fmt.Println("hello")
	var v int = 700

	if v < 1000 {
		fmt.Println("v is less than 1000")
	}
	var x int = 1
	if x >= 2 {
		x++
	} else {
		fmt.Print("In else")
	}

	fmt.Printf("\nValue of v is : %d\n", v)

	var v1 int = 4
	var v2 int = 7

	if v1 == 4 {
		if v2 == 7 {
			fmt.Printf("Value of v1 is 4 and v2 is 7\n")
		}
	}

	var str string = "Hey"
	if str == "Hey!" {
		fmt.Println("str is", str)
	} else if len(str) == 3 {
		fmt.Println("In else if length is 3")
	} else if len(str) > 3 {
		fmt.Println("In else if length is more than 3")
	} else {
		fmt.Println("In else")
	}

	/*
		This is wrong because there should be a curly brace before elseif like above
			if str=="Hey!" {
				fmt.Println("str is", str)
			}
			else if len(str)==3 {
				fmt.Println("In else if length is 3")
			}

		This too
			if x>1
			{
				do blah blah!
			}
	*/

	// Switch
	// Switch statement with both
	// optional statement, i.e, day:=4
	// and expression, i.e, day
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

	var val string = "five"
	switch val {
	case "one":
		fmt.Println("C#")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Java")
	}

	var value interface{}
	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)
	}
}
