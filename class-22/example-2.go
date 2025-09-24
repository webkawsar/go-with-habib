package main
import "fmt"

// Function Types
// Anonymous Function and IIFE
// A Noob Function Expression Example
// Function Expression or function assign in a variable


var a = 10

func sum() {
	add(3, 7)
}

func add(a int, b int) {
	fmt.Println(a+b)
}

func main() {
	fmt.Println("This text print from main")

	sum()
}

func init() {
	// example
	fmt.Println("This text print from init")
}

