package main
import "fmt"

// Function Types
// Init Function

var a = 10

func main() {
	fmt.Println("This text print from main")

	fmt.Println(a)
}

func init() {
	// 1st example
	fmt.Println("This text print from init")

	// 1st example
	fmt.Println(a)
	a = 20
	fmt.Println(a)
}


// at first e global scope e variable er value assign hoy
// then init function bose
// then main function
// jokhon dekhe init function ase tokhon aage init function execute kore dey
// then main function execute kore. ar main function execute hoye gele tokhon kaj ses hoye jay
