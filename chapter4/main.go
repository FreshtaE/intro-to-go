package main

import "fmt"

// func main() {
// 	// var x string
// 	// x = "first"
// 	// fmt.Println(x)
// 	// x = "second"
// 	// fmt.Println(x)

// 	x := "Hello Word"
// 	fmt.Println(x)

// 	f := "Freshta"
// 	fmt.Println("My name is", f)
// }

//Go is lexically scoped using blocks. Basically this means that the variable exists within the nearest curly braces { } (a block) including any nested curly braces (blocks), but not outside of them.

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	//Scanf is an inbuilt function from fmt package that reads user input
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println("Your number doubled is:", output)
}
