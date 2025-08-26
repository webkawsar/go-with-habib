
package main
import "fmt"

// scope
// global scope
// local scope
// block scope
// package scope


var p = 10
var q = 70

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func main() {

	// local scope
	// if, function, switch
	var a = 20
	var b = 30

	x := 100

	
	if x >= 50 {
		y := 500
		fmt.Println("Ami block er maje asi")
	}

	fmt.Println("Ami block er bahire asi", y)
}

