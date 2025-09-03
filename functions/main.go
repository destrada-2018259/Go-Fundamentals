package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func greeting(name string) string {

	return "Hello " + name
}

func main() {

	result := add(3, 5)

	fmt.Println("Result:", result)

	greeting := greeting("Daniel")

	fmt.Println(greeting)

}
