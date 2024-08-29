package main

import "fmt"
func double(x int) int{
	return x + x
}
func add(lhs,rhs int) int{
	return lhs + rhs
}
func greet(){
	fmt.Println("hello")
}
func main() {
   greet()
   dozen := double(6)
   fmt.Println("Dozen:",dozen)

   bakersDozen := add(dozen,1)
   fmt.Println("BakerDozen:",bakersDozen)

   another := add(double(6),1)
   fmt.Println("result",another)

}
