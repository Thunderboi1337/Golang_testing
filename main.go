package main

import "fmt"

func main() {

	// Playing with the basic syntax. Trying to understand. How I can write stuff hehe. Even writing this comment is a test.

	fmt.Println("Hello, World!")

	var nameOne string = "Jens"
	var nameTwo = "dude"

	fmt.Println(nameOne, nameTwo)
	fmt.Println(nameTwo)

	var numberOne int8 = 124
	var numberTWo = 255

	fmt.Println(numberOne)
	fmt.Println(numberTWo)

	var numbs float32 = 20.3

	fmt.Printf("\n Why v? %v. Well %d doesn't work.\n Bro %s doesn't want to join ", numbs, numberOne, nameOne)

	display_age()

}

//I think I already like Golang, feels like using a cousin of C/C++
