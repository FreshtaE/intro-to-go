//Package declaration - every Go program must start with a package declaration. Packages are Go's way of organizing and reusing code.
package main

//The import keyword is how we include code from other packages to use with our program. The fmt package (shorthand for format) implements formatting for input and output.
import "fmt"

//Function declaration - functions have inputs, outputs and a series of steps called statements which are executed in order. All functions start with the keyword func followed by the name of the function. This function has no parameters, doesn't return anything and has only one statement. The name main is special because it's the function that gets called when you execute the program.
func main() {
	//declare a variable
	var x string = "Hello World"
	fmt.Println(x)
	//Println stands for print line
	fmt.Println("Hello World")
	//finding the length of string
	fmt.Println(len("Hello World"))
	//accessing individual character in string
	fmt.Println("Hello World"[1])
	//concatenating two strings together
	fmt.Println("Hello" + "World")
}
