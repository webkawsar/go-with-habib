
package main
import "fmt"

// scope
// global scope


var p = 10
var q = 70

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func main() {
	var a = 20
	var b = 30

	
	fmt.Println(add(a, b))
	fmt.Println(add(p, q))
	// fmt.Println(add(p, sum))
}

