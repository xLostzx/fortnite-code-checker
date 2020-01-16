/*
  █████▒ ███▄    █  ▄████▄   ██░ ██ ▓█████  ▄████▄   ██ ▄█▀
▓██   ▒  ██ ▀█   █ ▒██▀ ▀█  ▓██░ ██▒▓█   ▀ ▒██▀ ▀█   ██▄█▒
▒████ ░ ▓██  ▀█ ██▒▒▓█    ▄ ▒██▀▀██░▒███   ▒▓█    ▄ ▓███▄░
░▓█▒  ░ ▓██▒  ▐▌██▒▒▓▓▄ ▄██▒░▓█ ░██ ▒▓█  ▄ ▒▓▓▄ ▄██▒▓██ █▄
░▒█░    ▒██░   ▓██░▒ ▓███▀ ░░▓█▒░██▓░▒████▒▒ ▓███▀ ░▒██▒ █▄
 ▒ ░    ░ ▒░   ▒ ▒ ░ ░▒ ▒  ░ ▒ ░░▒░▒░░ ▒░ ░░ ░▒ ▒  ░▒ ▒▒ ▓▒
 ░      ░ ░░   ░ ▒░  ░  ▒    ▒ ░▒░ ░ ░ ░  ░  ░  ▒   ░ ░▒ ▒░
 ░ ░       ░   ░ ░ ░         ░  ░░ ░   ░   ░        ░ ░░ ░
                 ░ ░ ░       ░  ░  ░   ░  ░░ ░      ░  ░
				   ░                       ░
												   by zSnails

Follow me on tiwtter https://twitter.com/zSnails
*/

package main

import (
	"flag"
	"fmt"

	log "../../logging"
)

var (
	//TxtFile the txt file
	TxtFile string
	//FnCode a single code to check
	FnCode string
	//CkURL djfh
	CkURL string
)

var (
	line1  = "	  █████▒ ███▄    █  ▄████▄   ██░ ██ ▓█████  ▄████▄   ██ ▄█▀\n"
	line2  = "	▓██   ▒  ██ ▀█   █ ▒██▀ ▀█  ▓██░ ██▒▓█   ▀ ▒██▀ ▀█   ██▄█▒\n"
	line3  = "	▒████ ░ ▓██  ▀█ ██▒▒▓█    ▄ ▒██▀▀██░▒███   ▒▓█    ▄ ▓███▄░\n"
	line4  = "	░▓█▒  ░ ▓██▒  ▐▌██▒▒▓▓▄ ▄██▒░▓█ ░██ ▒▓█  ▄ ▒▓▓▄ ▄██▒▓██ █▄\n"
	line5  = "	░▒█░    ▒██░   ▓██░▒ ▓███▀ ░░▓█▒░██▓░▒████▒▒ ▓███▀ ░▒██▒ █▄\n"
	line6  = "	 ▒ ░    ░ ▒░   ▒ ▒ ░ ░▒ ▒  ░ ▒ ░░▒░▒░░ ▒░ ░░ ░▒ ▒  ░▒ ▒▒ ▓▒\n"
	line7  = "	 ░      ░ ░░   ░ ▒░  ░  ▒    ▒ ░▒░ ░ ░ ░  ░  ░  ▒   ░ ░▒ ▒░\n"
	line8  = "	 ░ ░       ░   ░ ░ ░         ░  ░░ ░   ░   ░        ░ ░░ ░\n"
	line9  = "	                 ░ ░ ░       ░  ░  ░   ░  ░░ ░      ░  ░\n"
	line10 = "	                   ░                       ░\n"
)

var logger = log.Logger{Name: "CHECKER", Format: "[{{type}}@{{name}}]: {{message}}", Level: "debug", Print: true}

func init() {
	//logo
	logger.Raw("\n\n")
	logger.Raw(line1 + line2 + line3 + line4 + line5 + line6 + line7 + line8 + line9 + line10)
	logger.Raw("				by zSnails ~~")
	//end of logo
	flag.StringVar(&TxtFile, "f", "", "The file containing each code")
	flag.StringVar(&FnCode, "c", "", "Check a single code")
	flag.Parse()
}

func main() {
	for {
		fmt.Printf("help for help> ")
		checkInput()
	}

}
