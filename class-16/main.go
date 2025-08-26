package main
import (
	"fmt"
	"custom-package/mathlib"
)

// scope
// global scope
// local scope
// block scope
// package scope

var p = 10
var q = 70

func main() {

	// add(10, 15)

	// show a message
	fmt.Println("Showing here are result from custom package math.go file")

	// mathlib.multiply(4, 20)

	mathlib.Divide(30, 3)

	fmt.Println(mathlib.Message)
}


// notes
// 2ta file eksathe run kora jonno go run main.go add.go
// custom package pawar jonno
// go mod init example.com
// function er name small letter e dile function pabe na, capital letter diye function er name suru hote hobe
// jei name diye mod init korob seta import korar somoy package er namer aage bosate hobe







