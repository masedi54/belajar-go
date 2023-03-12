package main

import "fmt"

func main() {

	/*
		// basic map
		var person map[string]string     // Deklarasi map dengan tipe data string
		person = make(map[string]string) // Inisialisasi map

		person["name"] = "John Wick"
		person["age"] = "30"
		person["address"] = "Indonesia"

		fmt.Println("name:", person["name"])
		fmt.Println("age:", person["age"])
		fmt.Println("address:", person["address"])
	*/

	/*
		// looping with map
		var person = map[string]string{
			"name":    "John Wick",
			"age":     "30",
			"address": "Indonesia",
		}

		for key, value := range person {
			fmt.Println(key, ":", value)
		}
	*/

	/*
		// deleting Value
		var person = map[string]string{
			"name":    "John Wick",
			"age":     "30",
			"address": "Indonesia",
		}

		fmt.Println("Before deleting:", person)

		delete(person, "age")

		fmt.Println("After deleting:", person)
	*/

	/*
		// detecting a value
		var person = map[string]string{
			"name":    "John Wick",
			"age":     "30",
			"address": "Indonesia",
		}

		value, exist := person["age"]

		if exist {
			fmt.Println(value)
		} else {
			fmt.Println("Value does'nt exist")
		}

		delete(person, "age")

		value, exist = person["age"]

		if exist {
			fmt.Println(value)
		} else {
			fmt.Println("Value has been deleted")
		}
	*/

	// combining slice with map
	var people = []map[string]string{
		{"name": "John Wick", "age": "30"},
		{"name": "Ethan Hunt", "age": "32"},
		{"name": "Bourne", "age": "33"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
