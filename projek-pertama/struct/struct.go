/*
// giving values to struct fields
package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
	Divisi string
}

func main() {
	var employee Employee
	employee.Name = "John Doe"
	employee.Salary = 2000000
	employee.Divisi = "IT"

	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
	fmt.Println(employee.Divisi)
}
*/

/*
// initialize struct
package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
	Divisi string
}

func main() {
	var employee1 Employee
	employee1.Name = "John Doe"
	employee1.Salary = 2000000
	employee1.Divisi = "IT"

	var employee2 = Employee{Name: "John Bics", Salary: 3000000, Divisi: "UI/UX"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}
*/

/*
// pointer to struct
package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	Name   string
	Salary int
	Divisi string
}

func main() {
	var employee1 = Employee{Name: "John Doe", Salary: 2000000, Divisi: "IT"}
	var employee2 *Employee = &employee1

	fmt.Println("Employee1 name:", employee1.Name)
	fmt.Println("Employee2 name:", employee2.Name)

	fmt.Println(strings.Repeat("#", 20))

	employee2.Name = "John Bics"
	fmt.Println("Employee1 name:", employee1.Name)
	fmt.Println("Employee2 name:", employee2.Name)
}
*/

/*
// embedded struct
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Division string
	Person   Person
}

func main() {
	var employee1 = Employee{}

	employee1.Person.Name = "John Doe"
	employee1.Person.Age = 30
	employee1.Division = "IT"

	fmt.Printf("%+v", employee1)
}
*/

/*
// anonymous struct
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Anonymous struct tanpa pengisian field
	var employee1 = struct {
		Person   Person
		Division string
	}{}
	employee1.Person = Person{Name: "John Doe", Age: 30}
	employee1.Division = "IT"

	// Anonymous struct dengan pengisian field
	var employee2 = struct {
		Person   Person
		Division string
	}{
		Person:   Person{Name: "John Bics", Age: 30},
		Division: "UI/UX",
	}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}
*/

/*
// slice of struct
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	var people = []Person{
		{Name: "John Doe", Age: 30},
		{Name: "John Bics", Age: 30},
		{Name: "John Smith", Age: 30},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
*/

// slice of anonymous struct
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var employees = []struct {
		Person   Person
		Division string
	}{
		{Person: Person{Name: "John Doe", Age: 30}, Division: "IT"},
		{Person: Person{Name: "John Bics", Age: 30}, Division: "UI/UX"},
		{Person: Person{Name: "John Smith", Age: 30}, Division: "IT"},
	}

	for _, v := range employees {
		fmt.Printf("%+v\n", v)
	}
}
