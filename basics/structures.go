package basics

import (
	"fmt"
)

func RunStructures() {
	fmt.Println("Welcome! to structs!!")

	/*
		A structure or struct in Golang is a user-defined type that allows
		to group/combine items of possibly different types into a single type.

		It can be termed as a lightweight class that does not
		support inheritance but supports composition.

		type Address struct {
			name string
			street string
			city string
			state string
			Pincode int
		}

		To Define a structure:
			var a Address

		For above a all fields are default like string -> "" and int -> 0

		var a = Address{"Akshay", "PremNagar", "Dehradun", "Uttarakhand", 252636}
		Always pass the field values in the same order in which they are declared in the struct.

		To access individual fields of a struct you have to use dot (.) operator.
	*/

	member1 := member{
		name:   "John Doe",
		skills: []string{"Go", "Java", "Python"},
	}

	// Initialize another member
	member2 := member{
		name:   "Jane Smith",
		skills: []string{"C++", "JavaScript", "Ruby"},
	}

	// Create a team and populate it with members
	myTeam := team{
		members: []member{member1, member2},
		size:    2,
		lead:    "John Doe",
	}

	/*
		In Go, when you have a pointer to a struct, you can use the . operator
		directly on the pointer to access its fields. You don't need to dereference
		the pointer explicitly with *ptr. The Go compiler automatically handles
		the dereferencing when you use the . operator on a pointer.
		So, both ptr.members and (*ptr).members are equivalent in Go. Here's why:

		var ptr *team = &myTeam
		fmt.Println(ptr.members)    // Directly accessing the field using the pointer
		fmt.Println((*ptr).members) // Explicit dereferencing and then accessing the field

		In the first line, ptr.members, the Go compiler automatically dereferences
		the pointer ptr and then accesses the members field.
	*/
	var ptr *team = &myTeam
	fmt.Println(ptr.members, (*ptr).members)

	fmt.Println("Team Information:")
	fmt.Printf("Lead: %s\n", myTeam.lead)
	fmt.Printf("Team Size: %d\n", myTeam.size)

	//Anonymous Structure

	/*
		In Go language, you are allowed to create an anonymous structure.
		An anonymous structure is a structure which does not contain a name.
		It useful when you want to create a one-time usable structure.

		Syntex:
			variable_name := struct{
			// fields
			}{// Field_values}
	*/

	tommy := struct {
		name   string
		bark   bool
		weight int
	}{
		name:   "Tommy The Homie",
		weight: 50,
	}

	//By default bark will be false.
	fmt.Println("Hey dog's name is : ", tommy.name)

	//Anonymous Fields
	/*
		syntax:
			type struct_name struct{
				int
				bool
				float64
			}

		Invalid(can't have same type of fields) -> type student struct{
						int
						int
					}
	*/
	value := student{123, "Bud", 8900.23}

	fmt.Println("Enrollment number : ", value.int)
	fmt.Println("Student name : ", value.string)
	fmt.Println("Package price : ", value.float64)
}

type student struct {
	int
	string
	float64
}

type member struct {
	name   string
	skills []string
}

type team struct {
	members []member
	size    int
	lead    string
}
