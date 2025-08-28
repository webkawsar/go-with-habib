package main
import "fmt"

// Scope With Another Boring Example


var p = 10
var q = 70

func printNum(num int) {
	fmt.Println(num)
}

func add(a int, b int) {
	res := a + b
	printNum(res)
}

func main() {
	
	add(4, 6)
}






