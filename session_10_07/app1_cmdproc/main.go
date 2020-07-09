package main

import (
	"fmt"
	"os"
	"softdev_course/softdev_course/session_10_07/app1_cmdproc/commandprocessor"
	"strings"
	"time"
)

func main() {
	// commands
	// quit
	// getdate fmt
	// getdrives
	// listdir <dir>
	// hostname
	// sysinfo

	proc := &commandprocessor.CommandProcessor{}

	proc.Handle("quit", func(string, []string) bool {
		return false
	})

	proc.Handle("hostname", func(string, []string) bool {
		fmt.Printf("Host name is: %s\n", mustGetHostName())
		return true
	})

	proc.Handle("getdate", getdateHandler)

	proc.HandleDefault(defaultHandler)

	proc.Run()
}

func mustGetHostName() string {
	name, err := os.Hostname()

	if err != nil {
		panic(fmt.Sprintf("unable to get host name: %s", err))
	}

	return name
}

func getdateHandler(cmd string, args []string) bool {
	if len(args) == 0 {
		fmt.Println("Missing format parameter")
		return true
	}

	now := time.Now()

	var dateText string

	switch args[0] {
	case "date":
		dateText = now.Format("2006-01-02")
	case "time":
		dateText = now.Format("15:04:05")
	case "datetime":
		dateText = now.Format("2006-01-02 15:04:05")
	}

	if dateText != "" {
		fmt.Printf("Current date is: %s\n", dateText)
	} else {
		fmt.Printf("unrecognized format: %q\n", args[0])
	}

	return true
}

func defaultHandler(cmd string, args []string) bool {
	argsText := strings.Join(args, " ")
	fmt.Printf("unrecognized command: %s, [%s]\n", cmd, argsText)

	return true
}
