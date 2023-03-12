package main

/*
func main() {
	// deklarasi variable dengan tipe data uint8
	var a uint8 = 10
	var b byte // byte adalah alias dari uint8

	b = a // no error, karena byte memiliki tipe data yang sama dengan tipe data uint8
	_ = b
}
*/

import "fmt"

func main() {
	// mendeklarasikan tipe data alias bernama second
	// type nama_alias = nama_tipe_data

	type second = int

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint
}
