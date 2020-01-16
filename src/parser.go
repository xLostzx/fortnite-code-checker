package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func parseInput(input string) {
	var osSplitter string
	if runtime.GOOS == "windows" {
		osSplitter = "\\"
	} else {
		osSplitter = "/"
	}
	dir, err := os.Getwd()
	if err != nil {
		logger.Error("%v", err)
	}
	inputCode := strings.SplitN(input, ":", -1)
	if string(inputCode[0]) == "file" {
		filePath := (dir + osSplitter + inputCode[1])
		data, err := ioutil.ReadFile(filePath)
		println(string(data))
		if err != nil {
			logger.Critical("Either the file doesn't exist or you didn't specify one.")
		}
		s := string(data)
		arr := strings.SplitN(s, "\n", -1)
		checkCodes(arr)
		return
	}

	if string(inputCode[0]) == "code" {
		checkCode(string(inputCode[1]))
		return
	}

	switch input {
	case "help":
		logger.Raw(helpText)
		return
	case "exit", "done":
		os.Exit(1)
		return
	default:
		fmt.Printf("\n")
		fmt.Printf("fnchecker> ")
		return
	}
}

func checkInput() {
	var input string
	fmt.Scanf("%v", &input)
	parseInput(input)
}
