package main

/*
// slice cap function
import (
	"fmt"
	"strings"
)
*/

import "fmt"

func main() {
	/*
		// make function
		var fruits = make([]string, 3)

		_ = fruits

		fmt.Printf("%#v", fruits)
	*/

	/*
		// append function
		var fruits = make([]string, 3)

		fruits = append(fruits, "apple", "grape", "banana")

		fmt.Printf("%#v", fruits)
	*/

	/*
		// append function with ellipsis
		var fruits1 = []string{"apple", "grape", "banana"}

		var fruits2 = []string{"melon", "watermelon"}

		fruits1 = append(fruits1, fruits2...)

		fmt.Printf("%#v", fruits1)
	*/

	/*
		// slice copy function
		var fruits1 = []string{"apple", "grape", "banana"}
		var fruits2 = []string{"melon", "watermelon", "coklat"}

		nn := copy(fruits1, fruits2)

		fmt.Println("Fruits1 =>", fruits1)
		fmt.Println("Fruits2 =>", fruits2)
		fmt.Println("Copied element =>", nn)
	*/

	/*
		// slicing
		var fruits1 = []string{"apple", "grape", "banana", "melon", "watermelon", "coklat"}

		var fruits2 = fruits1[1:4]
		fmt.Printf("%#v\n", fruits2)

		var fruits3 = fruits1[0:]
		fmt.Printf("%#v\n", fruits3)

		var fruits4 = fruits1[:3]
		fmt.Printf("%#v\n", fruits4)

		var fruits5 = fruits1[:] // sama dengan fruits1[:len(fruits1)]
		fmt.Printf("%#v\n", fruits5)
	*/

	/*
		// combining slice and append
		var fruits1 = []string{"apple", "grape", "banana", "melon", "watermelon", "coklat"}

		fruits1 = append(fruits1[:3], "rambutan")

		fmt.Printf("%#v\n", fruits1)
	*/

	/*
		// slice backing array
		var fruits1 = []string{"apple", "grape", "banana", "melon", "watermelon", "coklat"}
		var fruits2 = fruits1[2:4]

		fruits2[0] = "orange"

		fmt.Println("fruits1 => ", fruits1)
		fmt.Println("fruits2 => ", fruits2)
	*/

	/*
		// slice cap function
		var fruits1 = []string{"apple", "grape", "banana", "melon"}

		fmt.Println("Fruits1 cap:", cap(fruits1)) //4
		fmt.Println("Fruits1 len:", len(fruits1)) //4

		fmt.Println(strings.Repeat("#", 20))

		var fruits2 = fruits1[0:3]

		fmt.Println("Fruits2 cap:", cap(fruits2)) //4
		fmt.Println("Fruits2 len:", len(fruits2)) //3

		fmt.Println(strings.Repeat("#", 20))

		var fruits3 = fruits1[1:]

		fmt.Println("Fruits3 cap:", cap(fruits3)) //3
		fmt.Println("Fruits3 len:", len(fruits3)) //3
	*/

	// creating new backing array
	cars := []string{"Ford", "Honda", "Toyota", "BMW", "Mercedes"}
	newCars := []string{}

	newCars = append(newCars, cars[1:3]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)
}
