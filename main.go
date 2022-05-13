package main

import (
	"fmt"
	"os"

	"github.com/website-monitor/ui"
	"github.com/website-monitor/website"
)

func main() {
	ui.LoadWelcomeMessage("")

	for {
		ui.LoadOptionsMenu()

		option := readCommand()

		switch option {
		case 1:
			ui.DisplayMessage("Starting monitoring. Please wait.")
			website.StartMonitoring()
		case 2:
			ui.DisplayMessage("Loading logs. Please wait.")
			website.DisplayLogs()
		case 0:
			ui.DisplayMessage("Thank you. Bye!")
			os.Exit(0)
		default:
			ui.DisplayMessage("I don't know this option. Please, rerun the program.")
			os.Exit(-1)
		}
	}
}

func readCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}
