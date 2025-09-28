package main
import "fmt"

// parameter vs argument
// argument mane aage
// parameter mane pore


func add(a int, b int) { // a,b parameter
	fmt.Println(a+b) // 13
}

func main() {
	fmt.Println("Hello World")

	add(4, 9) // 4,9 argument
}

