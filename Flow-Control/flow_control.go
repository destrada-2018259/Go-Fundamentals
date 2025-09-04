package main

import "fmt"

func main() {

	//If / else
	age := 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// If with initialization
	if n := 5; n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	//switch case

	day := "sunday"

	switch day {
	case "monday":
		fmt.Println("Start of the week.")

	case "friday":
		fmt.Println("Almost the end of the week.")

	default:
		fmt.Println("Any other day.")
	}

	//for

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While style

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	numbers := []int{10, 20, 30}

	for i, v := range numbers {
		fmt.Println("idex: ", i, " Value: ", v)
	}

	//User interaction

	var num int

	fmt.Println("please, give me a number")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Your number is even.")
	} else {
		fmt.Println("Your number is odd.")
	}

	//Numer comparison

	var num1 int
	var num2 int

	fmt.Println("Please give me a number")
	fmt.Scan(&num1)

	fmt.Println("Give me another one")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println(num1, " is bigger than ", num2)
	} else if num2 > num1 {
		fmt.Println(num2, " is bigger than ", num1)
	} else {
		fmt.Println(num2, " is equal to ", num1)
	}

	// Multiplication tables

	var table int

	fmt.Println("Enter the number you want to see its multiplication table")

	fmt.Scan(&table)

	for i := 1; i <= 10; i++ {

		fmt.Println(table, "x", i, "=", table*i)

	}

}
