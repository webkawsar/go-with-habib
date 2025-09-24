package main
import "fmt"

// Function Types
// Anonymous Function and IIFE
// A Noob Function Expression Example
// Function Expression or function assign in a variable


var a = 10

func main() {
	fmt.Println("This text print from main")

	add := func (a int, b int) {
		fmt.Println(a+b)
	}

	add(3, 5)
}

func init() {
	// example
	fmt.Println("This text print from init")
}

