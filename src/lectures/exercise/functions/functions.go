//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greeting(person string) {
	fmt.Println("Hello",person);
}

func msg() string{
	return "Hello"
}

func add(a int,b int,c int) int{
	return a + b + c
}

func number() int{
	return 5
}
func numbers() (int,int){
	return 5,6
}

func main() {
   greeting("Alice")
   fmt.Println(msg())
   result := add(number(),3,4)
   fmt.Println(result)
   number()
   numbers()

}
