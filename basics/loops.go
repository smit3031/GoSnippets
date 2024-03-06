package basics

import "fmt"

func RunLoops() int {

	fmt.Println("Welcome to looping!!!!!!")

	var i int = 2
	for ; i < 7; i++ {
		fmt.Printf("%d\n", i)
	}

	for i := 5; i > 1; i-- {
		fmt.Println(i)
	}

	/*
		A for loop is also used as an infinite loop by
		removing all the three expressions from the for loop.
		When the user did not write condition statement in for
		loop it means the condition statement is true and
		the loop goes into an infinite loop
	*/

	a := 2
	for {
		fmt.Println(a)
		if a > 5 {
			fmt.Printf("Inside if -> breaking this infinite loop. See you looping soon!\n")
			break
		}
		a++
	}

	/*
		you can pass or not pass this three and play around!
		for variableDeclare; condition; itr {

		}
	*/
	fmt.Println("Value of a is : ", a)
	b := &a

	for d := *b; true; {
		fmt.Println(d)
		if d == 0 {
			break
		}
		d -= 1
	}

	/*
		A for loop can also work as a while loop. This loop is executed until
		the given condition is true.
		When the value of the given condition is false the loop ends.
	*/
	x := 0
	for x <= 9 {
		fmt.Print(x)
		x++
	}
	fmt.Printf("\n")

	// Range loop
	rvariable := []string{"abc", "aabbcc", "aaabbbbccc"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j := range rvariable {
		fmt.Println(i, j)
	}

	/*
		A for loop can iterate over the key and value pairs of the map.
		Syntax:
		for key, value := range map {
			// Statement..
		}
	*/
	mmap := map[int]string{
		2: "Smit",
		1: "Yuvi",
		4: "Lucky",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}

	/*
		A for loop can iterate over the sequential
		values sent on the channel until it closed.
		Syntax:
			for item := range Chnl {
				// statements..
			}
	*/

	//Made a channel
	channel_name := make(chan int)

	//Anonymus function(invoked immd) to pass values to channel and close it.
	go func() {
		channel_name <- 100
		channel_name <- 1000
		channel_name <- 10000
		channel_name <- 100000
		close(channel_name)
		// channel_name <- 1
	}()

	for i := range channel_name {
		fmt.Println(i)
	}

	return 0
}
