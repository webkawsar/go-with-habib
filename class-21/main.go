package main
import "fmt"

// Function Types
// Anonymous Function and IIFE

var a = 10

func main() {
	fmt.Println("This text print from main")

	func (a int, b int) {
		fmt.Println(a+b)
	}(3, 8)
}

func init() {
	// example
	fmt.Println("This text print from init")
}

