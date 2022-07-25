package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now().Unix()

	argumentsAmount := len(os.Args[1:])

	if argumentsAmount > 0 {
		input := os.Args[1]
		if os.Args[1] == "help" || os.Args[1] == "-h" {
			printHelp()
			os.Exit(0)
		}

		if os.Args[1] == "version" || os.Args[1] == "-v" {
			printVersion()
			os.Exit(0)
		}

		if strings.HasPrefix(input, "-") || strings.HasPrefix(input, "+") {
			if strings.Count(input, "") < 3 {
				printError()
				os.Exit(1)
			}
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			finalTime := now + int64(inputInt)
			fmt.Println(finalTime)

		} else {
			if inputInt, err := strconv.Atoi(input); err == nil {
				var unixTime int64 = int64(inputInt)
				var strDate string

				t := time.Unix(unixTime, 0)

				if argumentsAmount == 2 && (os.Args[2] == "RFC3339" || os.Args[2] == "3339") {
					strDate = t.Format(time.RFC3339)
				} else {
					strDate = t.Format(time.UnixDate)
				}

				fmt.Println(strDate)
			} else {
				printError()
				os.Exit(1)
			}
		}
	} else {
		fmt.Println(now)

	}
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
	helpTxt += " 0.0.x\n\n"
	helpTxt += "https://github.com/iobear/uxt \n"

	fmt.Println(helpTxt)
}

func printError() {
	fmt.Println("Ivalid argument, try help argument")
}

func printVersion() {
	fmt.Println("0.0.1")

}
