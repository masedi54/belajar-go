package main

/*
// rune-by-rune way2
import (
	"fmt"
	"unicode/utf8"
)
*/

import "fmt"

func main() {

	/*
		// Looping over string (byte-by-byte)
		city := "Jakarta"

		for i := 0; i < len(city); i++ {
			fmt.Printf("index: %d, byte: %d\n", i, city[i])
		}
	*/

	/*
		// slice/slice of bytes
		var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}
		fmt.Println(string(city))
	*/

	/*
		// rune-by-rune
		str1 := "Hello"
		str2 := "World"

		fmt.Printf("str1 byte length => %d\n", len(str1))
		fmt.Printf("str2 byte length => %d\n", len(str2))
	*/

	/*
		// rune-by-rune way2
		str1 := "Hello"
		str2 := "World"

		fmt.Printf("str1 byte length => %d\n", utf8.RuneCountInString(str1))
		fmt.Printf("str2 byte length => %d\n", utf8.RuneCountInString(str2))
	*/

	// rune-by-rune way3
	str := "Hello World"

	for i, s := range str {
		fmt.Printf("index: %d, string: %s\n", i, string(s))
	}
}
