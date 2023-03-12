package main

import "fmt"

func main() {

	/*
		// menggunakan if else
		var currentYear int = 2022

		if age := currentYear - 2000; age < 17 {
			fmt.Println("You are not old enough to vote")
		} else {
			fmt.Println("You are old enough to vote")
		}
	*/

	/*
		// menggunakan switch
		var score = 8

		switch score {
		case 8:
			fmt.Println("Perfect")
		case 7:
			fmt.Println("Awesome")
		default:
			fmt.Println("Not Bad")
		}
	*/

	/*
		// switch with relational operator
		var score = 6

		switch {
			case score == 8:
				fmt.Println("Perfect")
				case (score < 8) && (score > 3):
					fmt.Println("not bad")
					default:
						{
							fmt.Println("study harder")
							fmt.Println("you need to learn more")
						}
		}
	*/

	/*
		// switch with fallthrough
		var score = 6

		switch {
		case score == 8:
			fmt.Println("Perfect")
			case (score < 8) && (score > 3):
				fmt.Println("not bad")
				fallthrough
			case score < 5:
				fmt.Println("It is ok, please study harder")
				default:
					{
						fmt.Println("study harder")
						fmt.Println("You don't have a good score yet")
					}
		}
	*/

	// nested condition
	var score = 18

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("Perfect")
			default:
				fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder")			
			}		
		}
	}
}