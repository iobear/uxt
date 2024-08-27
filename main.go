/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iobear/uxt/uxt"
)

var version = "v0.0.6"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(uxt.GetCurrentUnixTime())
		return
	}

	switch args[0] {
	case "help", "-h", "--help":
		printHelp()
	case "version", "-v", "--version":
		printVersion()
	default:
		processUnixTimeArgs(args)
	}
}

func processUnixTimeArgs(args []string) {
	input := args[0]

	if input == "since" {
		if len(args) < 2 {
			printError()
			return
		}
		unixTime, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			printError()
			return
		}

		fmt.Println(uxt.GetTimeSince(unixTime))
		return
	} else if input[0] == '+' || input[0] == '-' {
		adjustment, err := strconv.Atoi(input)
		if err != nil || len(input) < 2 {
			printError()
			return
		}

		result, err := uxt.AdjustCurrentUnixTime(adjustment)
		if err != nil {
			printError()
			return
		}

		fmt.Println(result)
	} else if strings.HasPrefix(input, "serv:") {
		port := strings.TrimPrefix(input, "serv:")
		RunServer(port)
		return
	} else {
		unixTime, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			printError()
			return
		}

		format := ""
		if len(args) == 2 {
			format = args[1]
		}

		result, err := uxt.ConvertUnixTimeToFormattedString(unixTime, format)
		if err != nil {
			printError()
			return
		}

		fmt.Println(result)
	}
}

func printError() {
	fmt.Println("Invalid argument, try help argument")
}

func printHelp() {
	helpTxt := "This is all the help you get  \n\n"
	helpTxt += os.Args[0] + "\n"
	helpTxt += " 946684800\n\n"
	helpTxt += os.Args[0] + " +<int> or -<int>\n"
	helpTxt += " 946684810\n\n"
	helpTxt += os.Args[0] + " <Unix Epoch> \n"
	helpTxt += " Thu May  8 08:13:30 CEST 2025\n\n"
	helpTxt += os.Args[0] + " <Unix Epoch> RFC3339\n"
	helpTxt += " 2025-05-08T08:13:30+02:00\n\n"
	helpTxt += os.Args[0] + " version\n"
	helpTxt += " " + version + "\n\n"
	helpTxt += "https://github.com/iobear/uxt \n"

	fmt.Println(helpTxt)
}

func printVersion() {
	fmt.Println(version)
}
