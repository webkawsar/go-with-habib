package main
import "fmt"

// parameter vs argument

func add(a int, b int) { // a,b parameter
	fmt.Println(a+b)
}

func main() {
	fmt.Println("Hello World")

	add(4, 9) // 4,9 argument
}

