/*
package main

func main() {
	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber
}
*/

/*
// memory address
package main

import "fmt"

func main() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("first number (value)		:", firstNumber)
	fmt.Println("first number (memori address)	:", &firstNumber)

	fmt.Println("second number (value)		:", *secondNumber)
	fmt.Println("second number (memori address)	:", secondNumber)
}
*/

/*
// pointer change value through pointer
package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value)		:", firstPerson)
	fmt.Println("firstPerson (memori address)	:", &firstPerson)
	fmt.Println("secondNumber (value)		:", *secondPerson)
	fmt.Println("secondNumber (memori address)	:", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value)		:", firstPerson)
	fmt.Println("firstPerson (memori address)	:", &firstPerson)
	fmt.Println("secondNumber (value)		:", *secondPerson)
	fmt.Println("secondNumber (memori address)	:", secondPerson)
}
*/

// pointer as function parameter
package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println("Before:", a)

	changeValue(&a)
	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
