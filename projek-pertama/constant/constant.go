package main 

import "fmt"

func main() {

	
	/*
		// const full_name string = "John Doe"
	
		fmt.Printf("Hello %s", full_name)
	*/

	/*
		// jika const tidak dikasi nilai maka akan menimbulkan error pada saat di compile

		const full_name string
		fmt.Println(full_name)
	*/

	/*
		// arithmatic operation
		var value = 2 + 2 * 3
		fmt.Println(value)

		// jika ingin penjumlahannya dieksekusi lebih awal maka menggunakan ()
		var value = (2 + 2) * 3
		fmt.Println(value)
	*/

	/*
		// relational operation
		var firstCondition bool = 2 < 3
		var secondCondition bool = "John"	== "John"
		var thirdCondition bool = 10 != 2.3
		var fourthCondition bool = 11<= 11

		fmt.Println("first condition", firstCondition)
		fmt.Println("second condition", secondCondition)
		fmt.Println("third condition", thirdCondition)
		fmt.Println("fourth condition", fourthCondition)
	*/

	// logical operation
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongRevers = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongRevers)
}