package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	ui "github.com/website-monitor/ui"
)

const monitorDelay = 5 * time.Second
const monitorQuantity = 2

func main() {
	ui.LoadWelcomeMessage("")

	for {
		ui.LoadOptionsMenu()

		option := readCommand()

		switch option {
		case 1:
			fmt.Println("Started monitoring...")
			siteMonitoring()
		case 2:
			fmt.Println("Diplaying logs")
			displayLogs()
		case 0:
			fmt.Println("Exit program")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
			os.Exit(-1)
		}
	}
}

func readCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}

func siteMonitoring() {
	sites := getSites()

	for i := 0; i < monitorQuantity; i++ {
		for _, site := range sites {
			if isSiteUp(site) {
				fmt.Println("The site ", site, " is up!")
				registerIntoLogs(site, true)
			} else {
				fmt.Sprintln("The site ", site, " is down!")
				registerIntoLogs(site, false)
			}
		}

		time.Sleep(monitorDelay)
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
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerIntoLogs(site string, status bool) {
	file, err := os.OpenFile("logs", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		displayErr(err)
	}

	logTime := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString("[time] " + logTime + " [site] " + site + " UP: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func displayLogs() {
	file, err := ioutil.ReadFile("logs")

	if err != nil {
		displayErr(err)
	}

	fmt.Println(string(file))
}

func displayErr(err error) {
	fmt.Println("An error has ocurred: ", err)
}
