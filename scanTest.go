package main

import (
	"./scanner/scanner"
	"./scanner/basic"
	"./scanner/table"
	"./../misc/timer"
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	t := timer.NewTimer()
	content := load("sample.txt")
	test(basic.NewScanner(content), t, "basic scan")
	test(table.NewScanner(content), t, "table scan")
}

func test(s scanner.Interface, t *timer.Timer, testName string) {
	t.Add(testName)
	err, ok := s.Scan()
	if !ok {
		fmt.Printf("Error: %s\n", err.Value)
		return
	}
	t.Get(testName).Stop()
	fmt.Printf("%s: %f seconds\n", testName, t.Get(testName).Value)
}

func print(s scanner.Interface, t *timer.Timer, testName string) {
	pos := 0
	tok := s.GetToken(pos)
	for ; tok != nil; tok = s.GetToken(pos) {
		fmt.Printf("%s [%s]\n", scanner.GetTokenName(tok.Type), tok.Value)
		pos++
	}
	fmt.Printf("Time taken: %f seconds\n", t.Get(testName).Value)
}

func load(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading from \"%s\": %s\n", filename, err.String())
		os.Exit(1)
	}
	return content[0:len(content)-1]	// chop off last char (unwanted newline)
}
