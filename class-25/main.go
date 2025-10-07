package main
import "fmt"

// go er code run hoy 2ta way te
// 2 ta Phase
// compilation
// execution


// compilation phase
// sokol constant bose
// function definition bose


const a = 10
var p = 120

func call() {
	add := func (x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}

func main() {
	call()

	fmt.Println(a)
}

func init() {
	fmt.Println("Hello from init")
}


/*

	******* Compilation Phase *******
	ei phase e code binary file e convert hoy

	** Code Segment **
	a = 10
	call = func () {....}
	add = func () {....}
	main = func () {....}
	init = func () {....}


	******** Execution Phase **********

	


	**** Binary File ****
	go build main.go
	./main

	go run main.go => compile it => main => ./main
	go build main.go => compile it => main
	./main

*/


// compilation phase e check korbe main, init function ache kina
// compilation phase ei binary file er majer code segment e ache
// je function ta run hoy tokhon stack er majhe koyekta cell niye seti stack frame hoy
