package main

import (
	sBasic "./scanner/basic"
	"./parser/parser"
	pBasic "./parser/basic"
	"./../misc/timer"
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	content := load("sample.txt")
	t := timer.NewTimer()
	test(pBasic.NewParser(sBasic.NewScanner(content)), t, "basic parse")
}

func test(p parser.Interface, t *timer.Timer, testName string) {
	t.Add(testName)
	err := p.Parse()
	if err.GetType() != parser.NO_ERROR {
		fmt.Printf("%s: %s, %s\n", err.GetType(), err.GetFunc(), err.GetDesc())
	}
	t.Get(testName).Stop()
	fmt.Printf("%s: %f seconds\n", testName, t.Get(testName).Value)
}

func load(filename string) []byte {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading from \"%s\": %s\n", filename, err.String())
		os.Exit(1)
	}
	return content[0:len(content)-1]	// chop off last char (unwanted newline)
}
