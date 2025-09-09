package main

import (
	"fmt"
)

func main() {

	//slice / dynamic

	numbers := []int{1, 2, 3, 4}

	numbers = append(numbers, 5)

	fmt.Println("This is a slice using Go: ", numbers)

	//Using maps

	users := map[string]string{
		"name": "Daniel",
		"age":  "20",
	}

	fmt.Println("User", users["name"], users["age"])

	ages := []int{20, 19, 20, 13, 21}

	for i := 0; i < len(ages); i++ {
		fmt.Println("Age: ", ages[i])
	}

}
