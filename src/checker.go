package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func checkCodes(codes []string) {
	for i := 0; i < len(codes); i++ {
		time.Sleep(time.Second)
		checkCode(codes[i])
	}
}

func checkCode(code string) {
	CkURL = "https://api.nitestats.com/v1/codes/checker?codes=" + string(code)
	resp, err := http.Get(CkURL)
	if err != nil {
		logger.Error("%v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("%v", err)
		return
	}
	s := string(body)
	s = strings.Replace(s, "<br>", "\n", 1)
	logger.Info(s)
}
