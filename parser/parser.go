// Package defines a recursive descent parser for a simple calculator language
package parser

import (
	"./../../misc/error"
	"./../scanner/scanner"
)

const NO_ERROR = "no error"

// ParseError implements error.Interface.
// It provides useful parse-specific information.
type ParseError struct {
	nonTerminal string		// refers to the non-terminal which this error violates
	expected []string		// expected tokens
	got string				// actual token
}

func NewError(nt string, exp []int, actual int) *ParseError {
	p := &ParseError{nt, []string{}, scanner.GetTokenName(actual)}
	for i := 0; i < len(exp); i++ {
		p.expected = append(p.expected, scanner.GetTokenName(exp[i]))
	}
	return p
}

func (p *ParseError) GetType() string {
	return "Parse error"
}

func (p *ParseError) GetDesc() string {
	desc := "Expected ["
	for i := 0; i < len(p.expected); i++ {
		desc += p.expected[i]
		if (i + 1) < len(p.expected) {
			desc += " | "
		}
	}
	desc += "] but got [" + p.got + "]"
	return desc
}

func (p *ParseError) GetFunc() string {
	return p.nonTerminal
}

type Interface interface {
	Parse() error.Interface
}
