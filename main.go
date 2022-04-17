package main

import (
	"fmt"
	"os"
)

func main() {
	displayWelcome()

	for {
		displayMenu()

		option := readCommand()

		switch option {
		case 1:
			fmt.Println("Option 1")
		case 2:
			fmt.Println("Option 2")
		case 0:
			fmt.Println("Exit program")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
			os.Exit(-1)
		}
	}
}

func displayWelcome() {
	name := "Fulano"
	version := 1.0

	fmt.Println("Hello, " + name + "!")
	fmt.Println("Program version: ", version)
}

func displayMenu() {
	fmt.Println("")
	fmt.Println("Choose an option:")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("0 - Signout")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}
