package main

/*
// basic and variadic functions
import "fmt"
*/

// return & variadic function #3
import (
	"fmt"
	"strings"
)

/*
// return multiple values
import (
	"fmt"
	"math"
)
*/

/*
func main() {
	great("Airell", 23)
}

func great(name string, age int) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}
*/

/*
func main() {
	great("Airell", "Jalan Raya")
}

func great(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}
*/

/*
// return
func main() {
	var names = []string{"Airell", "Ari"}
	var printMsg = greet("Hei", names)

	fmt.Println(printMsg)
}

func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}
*/

/*
// returning multiple values & predifened return values
func main() {
	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)
}

func calculate(d float64) (float64, float64) {
	// menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}
*/

/*
// variadic functions #1
func main() {
	studentLists := print("Airell", "Ari", "Ariq", "Ariqul", "Ariqulhaq")

	fmt.Printf("%v", studentLists)
}

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}
*/

/*
// variadic functions #2
func main() {
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := sum(numberLists...)

	fmt.Println("Result:", result)
}

func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}
	return total
}
*/

// variadic functions #3
func main() {
	profile("Airell", "Jalan Raya", "Pasta", "Pizza", "Burger", "Nasi Goreng")
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("My name is", name)
	fmt.Println("My favorite foods are", mergeFavFoods)
}
