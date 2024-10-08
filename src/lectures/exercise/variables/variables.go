//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:

//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favColor = "black"
	fmt.Println("fav color is",favColor)
    //* Store your birth year and age (in years) in two variables using
    //  compound assignment
     birthYear,age := 2002,22
	 fmt.Println("Born in",birthYear,"age =",age,"years")
//* Store your first & last initials in two variables using block assignment
     var(
		firstInitial = "S"
		lastInitial = "P"
	 )
	 fmt.Println("Initials=",firstInitial,lastInitial)
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
    var ageInDays int
	ageInDays = 365 * age
	fmt.Println("I'm",ageInDays,"days old")

    a,b := 6,3
	fmt.Println("Sum: ", a+b)
	fmt.Println("Diff: ", a-b)
	fmt.Println("Multiplication: ", a*b)
	fmt.Println("Divison: ", a/b)
}
