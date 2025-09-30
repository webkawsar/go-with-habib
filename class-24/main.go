package main
import "fmt"

// Internal memory | Code Segment | Data Segment | Stack | Heap | GC
// sokol global variable RAM er Data segment e store hoy
// sokol global function RAM er Code Segment e store hoy
// function call/invoke hole stack frame create hoy
// Heap ke manage kore GC or Garbage Collector


var a = 10

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	fmt.Println("This text print from main")

	add(4, 3)
	add(a, 7)
}

func init() {
	fmt.Println("This text print from init")
}

