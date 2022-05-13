package main

import (
	"fmt"
	"os"

	ui "github.com/website-monitor/ui"
	"github.com/website-monitor/website"
)

func main() {
	ui.LoadWelcomeMessage("")

	for {
		ui.LoadOptionsMenu()

		option := readCommand()

		switch option {
		case 1:
			fmt.Println("Starting monitoring. Please wait.")
			website.StartMonitoring()
		case 2:
			fmt.Println("Loading logs. Please wait.")
			website.DisplayLogs()
		case 0:
			fmt.Println("Thank you. Bye!")
			os.Exit(0)
		default:
			fmt.Println("I don't know this option. Please, rerun the program.")
			os.Exit(-1)
		}
	}
}

func readCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}
