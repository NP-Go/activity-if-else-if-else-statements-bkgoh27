package main

import "fmt"

func authenticateUser(name string) string {
	greeting := "Welcome"

	if name == "Admin" {
		greeting = greeting + ", " + name + "!"
	} else if name == "Robin" || name == "John" {
		greeting += "!"
	} else {
		greeting = "Merry Men"
	}
	fmt.Println(greeting)
	return greeting
}

func main() {
	var name string
	fmt.Println("Give me your name!")
	fmt.Scanln(&name)

	authenticateUser(name)
}
