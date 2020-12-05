package essential

import (
	"fmt"
)

// Here are some global variables
const (
	// Pi is just pi, but exported const (begin with uppercase letters) should have a comment
	// it is visable from outside the package (public)
	Pi float32 = 3.1415926
	e  float32 = 123.321 // This is a unexported const, and it is invisable from outside this package (private)
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

func declareArr() {
	a := [5]int{1, 2, 3} // Declare an array of length 5, initialized as [1,2,3,0,0]
	// Declare an array without length setup, go will compute it automatically. in this case
	// b will be an array of 5 ints
	b := [...]int{2, 3, 4, 5, 6}
	fmt.Println(a)
	fmt.Printf("The length of b is: %d\n", len(b))
}

func setFirstEleToOne(slc []int) {
	// Set the first element of a slice into 1
	slc[0] = 1
}

func sliceUsage() {
	// Slice is a reference of a given array, it is like the slice in python.
	// Note it is a reference, i.e., change the value in slice will change
	// the original array!
	var givenArr [10]int
	sliceInFirstFive := givenArr[:5]
	fmt.Println("Before the set: ", sliceInFirstFive)
	setFirstEleToOne(sliceInFirstFive)
	fmt.Println("After the set: ", sliceInFirstFive)
	fmt.Println("Array value after the set: ", givenArr)
	// A slice have two main method: cap() and len(), for its capacity and length
	// cap is its base array's length
	// length is the length of slice
	fmt.Println("The slice's cap is: ", cap(sliceInFirstFive), ", And the length is: ", len(sliceInFirstFive))
	// One can use three args to declare a slice:
	slc := givenArr[:5:8]
	// The omitting first one is 0 as default, and the last one is its capacity
	fmt.Println("The slice's cap is: ", cap(slc), ", And the length is: ", len(slc))

}

func setMapVal(x map[string]int) {
	// Set the value of key "set" in x as 1
	x["set"] = 1
}

func declareMap() {
	// map is similar to dictionary in python
	amap := map[string]int{"a": 1, "b": 2}
	fmt.Println("here is a map:", amap)
	fmt.Println("its value of key\"a\" is: ", amap["a"])
	// map is also a reference implementation, i.e., it enters a function will not be copied
	setMapVal(amap)
	fmt.Println("After the set: ", amap)

}

func newAndMake() {
	// new will initialize all elements and return a ptr:
	ptr := new([10]int)
	fmt.Println("The returned value of new: ", ptr, ", and its content: ", *ptr)
	// make returns the value with out initialization, but can only used for map, slice, channel
	mkmap := make(map[string]int)
	fmt.Println("The returned value of make(map): ", mkmap)
	mkslc := make([]int, 5, 10) // The second args is its length, the third is its capacity
	fmt.Println("The returned value of make(slice)", mkslc, ", and its lengths/ caps", len(mkslc), "/", cap(mkslc))
}

func init() {
	fmt.Println("hellow!")
	declareVals()
	declareArr()
	sliceUsage()
	declareMap()
	newAndMake()
	Greater(1, 0)
}
