package main 

/*
// array modify element through index
import (
	"fmt"
	"strings"
)
*/

// array multidimensi
import (
	"fmt"
)

func main() {

	/*
	// array first way
	var numbers [4] int
	numbers = [4] int {1, 2, 3, 4}

	var strings = [3] string {"a", "b", "c"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)
	*/

	/*
	// array modify element through index
	var fruits = [3] string {"apple", "banana", "grape"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "grape"

	fmt.Printf("%#v\n", fruits)
	*/

	/*

	// array loop through elements
	var fruits = [3] string {"apple", "banana", "grape"}

	// cara pertama
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	// cara kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	}
	*/

	// array multi dimensi
	balances := [2][3]int{{ 1, 2, 3 }, { 4, 5, 6 }}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}