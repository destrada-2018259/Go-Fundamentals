package main

import "fmt"

func main() {

	// Explicit Declaration
	var name string = "Daniel"

	//Type inference
	// := Declares a new Variable

	// name := "Daniel" == var name string = "Daniel"
	age := 20

	//constants
	const favoriteManga = "Chainsawman"

	fmt.Println("Your Data: ", name, age, favoriteManga)
}
