
package main
import "fmt"

// introduction to functions
// Function with Return Values and Types
// multiple jinis return korar way
func add(num1 int, num2 int) (int,int) {
	sum := num1 + num2

	mul := num1 * num2

	return sum, mul
}

func main() {
	a := 10
	b := 20

	sum, mul := add(a, b)
	fmt.Println(sum)
	fmt.Println(mul)
}

