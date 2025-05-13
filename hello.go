package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// string variable
	var name string
	name = "John Doe"
	fmt.Println("Hello, " + name + "!")

	lastname := "Brahmin"
	fmt.Println("Latname: " + lastname)

	fmt.Printf("Name: %s, Lastname: %s\n", name, lastname)

	fmt.Printf(("Salary: %f\n"), 1000.23)

	fmt.Printf("Salary: %.2f\n", 1000.238)

	fmt.Printf("Age: %d\n", 25)

}
