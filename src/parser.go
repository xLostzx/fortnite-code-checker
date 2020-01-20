package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func parseInput(input string) error {
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
		if len(inputCode) == 0 {
			logger.Error("Either the file doesn't exist or you didn't specify one.")
		}
		filePath := (dir + osSplitter + inputCode[1])
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			logger.Critical("Either the file doesn't exist or you didn't specify one.")
		}
		s := string(data)
		arr := strings.SplitN(s, "\n", -1)
		checkCodes(arr)
		return nil
	}

	if string(inputCode[0]) == "code" {
		checkCode(string(inputCode[1]))
		return nil
	}

	switch input {
	case "help":
		logger.Raw(helpText)
		break
	case "exit", "done":
		os.Exit(1)
		break
	default:
		fmt.Printf("\n")
		break
	}
	return nil
}

func checkInput() error {
	var input string
	fmt.Scanf("%v", &input)
	err := parseInput(input)
	if err != nil {
		return err
	}
	return nil
}
