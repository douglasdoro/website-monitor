package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	displayWelcome()

	for {
		displayMenu()

		option := readCommand()

		switch option {
		case 1:
			fmt.Println("Started monitoring...")
			siteMonitoring()
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

func siteMonitoring() {
	sites := getSites()

	for _, site := range sites {
		if isSiteUp(site) {
			fmt.Println("The site ", site, " is up!")
		} else {
			fmt.Sprintln("The site ", site, " is down!")
		}
	}
}

func isSiteUp(site string) bool {
	response, err := http.Get(site)

	if err != nil {
		displayErr(err)
	}

	if response.StatusCode == 200 {
		return true
	}

	return false
}

func getSites() []string {
	var sites []string

	file, err := os.Open("sites")

	if err != nil {
		displayErr(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println("Site: ", line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	return sites
}

func displayErr(err error) {
	fmt.Println("An error has ocurred: ", err)
}
