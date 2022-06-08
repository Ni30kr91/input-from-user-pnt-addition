// Golang program to show how
// to take input from the user
package main

import "fmt"

// main function
func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your First Number ")

	// var then variable name then variable type
	var first int 
	fmt.Scanln(&first)
	fmt.Println("number is", first)
	
	fmt.Println("Enter Your second Number ")
	// var then variable name then variable type
	var second int 
	fmt.Scanln(&second)
	fmt.Println("number is", second)

	var sum = first + second 
    fmt.Println("sum is",sum) 
	

}