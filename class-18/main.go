package main
import "fmt"

// Variable Shadowing

var a = 10

func main() {
	age := 30

	if age > 18 {
		a := 47
		fmt.Println(a) // 47
	}

	fmt.Println(a) // 10
}






