package main

import "fmt"

// Global Variable
var angka1 = 10

func main() {
	// Printing before declaring a local variable with the same name
	fmt.Println("before declaring a local variable : ", angka1)

	// Local Variable
	angka1 := 100

	// Printing after declaring a local variable with the same name
	fmt.Println("after declaring a local variable : ", angka1)

	{
		// just this scope Variable
		angka1 := 1000
		// Printing after declaring a Global variable and local variable with the same name
		fmt.Println(angka1)
	}

	fmt.Println("?", angka1) // 100
}
