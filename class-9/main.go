package main
import "fmt"

func main() {

	// if else condition

	age := 45
	if age < 18 {
		fmt.Println("you are a child")
	} else if age > 18 && age < 35 {
		fmt.Println("you are a Young Man")
	} else {
		fmt.Println("you are a Old Man")
	}


	// switch case
	salary := 200
	switch salary {
	case 100:
		fmt.Println("Salary is: 100")
	case 200, 300:
		fmt.Println("Salary is: 200 | 300")
	case 600:
		fmt.Println("Salary is: 600")
	default:
		fmt.Println("Salary is: 000")
	}
}