package main

import "fmt"

func main() {

	/*
	// first way of looping
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	*/

	/*
	// second way of looping
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}
	*/

	/*
	// third way of looping
	var i = 0

	for {
		fmt.Println("Angka", i)

		i++

		if i == 3 {
			break
		}
	}
	*/

	/*
	// looping break and continue keywords
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
	*/

	/*
	// nested looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	*/

	// label looping
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Outer Loop - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}

			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}