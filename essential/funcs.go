package essential

import (
	f "fmt"
)

// Tog is a string, public of this package
const Tog string = "This is module named pkg1"

// Greater will return a boolean value for whether a is greater than b
func Greater(a int, b int) bool {
	if a > b {
		f.Println("a is greater than b")
		return true
	}

	f.Println("a is smaller than b")
	return false
}

func printTri() {
	for line := 0; line < 6; line++ {
		for i := 0; i < line; i++ {
			f.Print("*")
		}
		f.Print("\n")
	}
}

func tryDefer(add int) {
	a := 3
	f.Printf("Now the init val is: %d", a)
	defer f.Printf("The fin val is: %d", a)
	a += add
}

func init() {

}
