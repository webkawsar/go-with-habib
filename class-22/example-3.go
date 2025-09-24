package main
import "fmt"


func main() {
	fmt.Println("This text print from main")

	calculate(3, 6)
	
	calculate := func (a int, b int) {
		fmt.Println(a+b)
	}

	calculate(2, 5)
}

func init() {
	// example
	fmt.Println("This text print from init")
}

