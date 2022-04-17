package main

import "fmt"

func main() {
	displayWelcome()
}

func displayWelcome() {
	name := "Fulano"
	version := 1.0

	fmt.Println("Hello, " + name + "!")
	fmt.Println("Program version: ", version)
}
