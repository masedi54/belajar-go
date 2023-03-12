package main

// basic and closure callback
import "fmt"

/*
// closure as a return value
import (
	"fmt"
	"strings"
)
*/

// func main() {

/*
	// declare closure in variable
	var eventNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(eventNumbers(numbers...))
*/

/*
	// closure IIFE
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)
*/

/*
	// closure as a return value
	var studentLists = []string{"John", "Wick", "Ethan", "Hunt", "Jack"}

	var find = findStudent(studentLists)

	fmt.Println(find("wick"))
}

func findStudent(students []string) func(string) string {

	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'nt exist", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position)
	}
}
*/

/*
// closure as a callback
func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
*/

// closure as a callback way2
type isOddNum func(int) bool

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}

func findOddNumbers(numbers []int, callback isOddNum) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
