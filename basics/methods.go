package basics

import "fmt"

func RunMethods() {
	fmt.Println("Welcome to go methods!")

	/*
		Go language support methods. Go methods are similar to Go function with
		one difference, i.e, the method contains a receiver argument in it.

		When you create a method in your code the receiver and receiver
		type must be present in the same package. And you are not allowed
		to create a method in which the receiver type is already defined in
		another package including inbuilt type like int, string, etc.
		If you try to do so, then the compiler will give an error.


		Syntax:

			func(reciver_name Type) method_name(parameter_list)(return_type){
			// Code
			}
	*/

	she := girl{
		name:  "Foram",
		crazy: 6,
		hot:   9,
	}

	she.showGirl()

	value1 := data(2)
	value2 := data(3)
	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)

	// Methods with Pointer Receiver
	/*
		Syntax:
			func (p *Type) method_name(...Type) Type {
			// Code
			}
	*/
	fmt.Println("Her hotness is : ", she.hot)
	she.incHot(1)
	fmt.Println("Her hotness now is : ", she.hot)

	/*
		Method Can Accept both Pointer and Value
		As we know that in Go, when a function has a value argument,
		then it will only accept the values of the parameter, and
		if you try to pass a pointer to a value function,
		then it will not accept and vice versa. But a Go method can
		accept both value and pointer, whether it is defined with pointer or value receiver.
	*/
	she.showName()
	(&she).changeName("Rachel")
	fmt.Println("Now her name is : ", she.name)

}

type girl struct {
	name  string
	crazy int
	hot   int
}

func (gf girl) showGirl() {
	fmt.Println("Her name is : ", gf.name)
	fmt.Println("Her craziness : ", gf.crazy)
	fmt.Println("Her hotness : ", gf.hot)
}

func (gf *girl) incHot(inc_hotness int) {
	(*gf).hot += inc_hotness
}

func (g *girl) changeName(new_name string) {
	(*g).name = new_name
}

func (g girl) showName() {
	g.name = "Phoebe"
	fmt.Println("Before her name was : ", g.name)
}

type data int

// Defining a method with non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}
