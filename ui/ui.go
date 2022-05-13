package ui

import "fmt"

func LoadWelcomeMessage(name string) {
	var welcomeMessage string = "Hello, admin! "

	if name != "" {
		welcomeMessage = "Hello, " + name + "!"
	}

	version := 1.0

	fmt.Println(welcomeMessage)
	fmt.Println("Program version: ", version)
}

func LoadOptionsMenu() {
	fmt.Println("")
	fmt.Println("Choose an option:")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("0 - Signout")
}

func DisplayErr(err error) {
	fmt.Println("Something is wrong. Please, check the error: ", err)
}
