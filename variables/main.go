package main

import "fmt"

// Here are some global variables
const (
	// Pi is just pi, but exported const (begin with uppercase letters) should have a comment
	// it is visable from outside the package
	Pi float32 = 3.1415926
	e  float32 = 123.321 // This is a unexported const, and it is invisable from outside this package
)

func declareVals() {
	var a int32     // Declare an int32 variable
	a = 32          // Set the variable as 32
	var b int32 = 4 // Declare an int32 variable and set it as 4
	var c = a + b   // Declare an vriable without type declaration
	d := c          // Short diclare an variable, can only used in a function
	_, e := d, 20   // Special _, it will be abandoned and cannot put into use
	fmt.Println("Just starting function \"declareVals()\"")
	fmt.Printf("a = %d, b = %d, c = %d, d = %d, e = %d, Pi = %.4f\n", a, b, c, d, e, Pi)
	fmt.Println("Exiting function \"declareVals()\"")
}

func main() {
	fmt.Println("hellow!")
	declareVals()
}
