// Package defines a simple switch-based scanner - *not* based on a DFA
package basic

import "./scanner"

type Scanner struct {
	Input []byte
	Tokens []*scanner.Token
	pos int
}

func NewScanner(input []byte) *Scanner {
	return &Scanner{input, []*scanner.Token{}, 0}
}

// Generates all tokens, returning on Error or End
func (bs *Scanner) Scan() (*scanner.Token, bool) {
	tok := bs.nextToken()
	for ; tok.Type != scanner.END; tok = bs.nextToken() {
		if tok.Type == scanner.ERROR {
			return tok, false
		}
		bs.Tokens = append(bs.Tokens, tok)
	}
	return nil, true
}

// Gets the next token from the input
func (bs *Scanner) nextToken() *scanner.Token {
	if bs.pos >= len(bs.Input) {
		return &scanner.Token{scanner.END, ""}
	}
	
	var char byte
	for ; bs.pos < len(bs.Input); bs.pos++ {
		char = bs.Input[bs.pos]
		if char != ' ' && char != '\n' && char != '\t' {
			break
		}
	}
	
	tok, ok := bs.getCharToken(char)
	if ok {
		bs.pos++
		return tok
	}
	
	switch char {
	
	case ':':
		bs.pos++
		if bs.Input[bs.pos] == '=' {
			bs.pos++
			return &scanner.Token{scanner.ASSIGN, ":="}
		}
		return &scanner.Token{scanner.ERROR, "symbol not recognised"}
		
	case '/':
		bs.pos++
		if bs.Input[bs.pos] == '*' {
			for ; bs.pos < len(bs.Input); bs.pos++ {
				if bs.Input[bs.pos] == '*' {
					bs.pos++
					if bs.pos >= len(bs.Input) {
						break
					}
					if bs.Input[bs.pos] == '/' {
						bs.pos++
						return bs.nextToken()
					}
				}
			}
			return &scanner.Token{scanner.ERROR, "block comment not escaped"}
		} else if bs.Input[bs.pos] == '/' {
			for ; bs.pos < len(bs.Input); bs.pos++ {
				if bs.Input[bs.pos] == '\n' {
					bs.pos++
					return bs.nextToken()
				}
			}
			return &scanner.Token{scanner.END, ""}
		}
		return &scanner.Token{scanner.DIV, "/"}
		
	case '.':
		bs.pos++
		if scanner.IsDigit(bs.Input[bs.pos]) {
			num := "0."
			for ; bs.pos < len(bs.Input) && scanner.IsDigit(bs.Input[bs.pos]); bs.pos++ {
				num += string(bs.Input[bs.pos])
			}
			return &scanner.Token{scanner.NUMBER, num}
		}
		return &scanner.Token{scanner.ERROR, "expected a number following '.'"}
	}
	
	if scanner.IsDigit(char) {
		points := 0
		var num string
		for ; bs.pos < len(bs.Input); bs.pos++ {
			if scanner.IsDigit(bs.Input[bs.pos]) {
				num += string(bs.Input[bs.pos])
			} else if bs.Input[bs.pos] == '.' {
				if points == 0 {
					num += string(bs.Input[bs.pos])
					points++
				} else {
					break
				}
			} else {
				break
			}
		}
		return &scanner.Token{scanner.NUMBER, num}
	}
	
	if scanner.IsLetter(char) {
		var value string
		for ; bs.pos < len(bs.Input); bs.pos++ {
			if value == "read" {
				return &scanner.Token{scanner.READ, value}
			} else if value == "write" {
				return &scanner.Token{scanner.WRITE, value}
			}
			if scanner.IsLetter(bs.Input[bs.pos]) {
				value += string(bs.Input[bs.pos])
			} else {
				break
			}
		}
		return &scanner.Token{scanner.ID, value}
	}
	
	bs.pos++
	return &scanner.Token{scanner.ERROR, "character '" + string(char) + "' not recognised"}
}

func (bs *Scanner) GetToken(i int) *scanner.Token {
	if i > -1 && i < len(bs.Tokens) {
		return bs.Tokens[i]
	}
	return nil
}

func (bs *Scanner) getCharToken(char byte) (tok *scanner.Token, ok bool) {
	ok = true
	t := scanner.ERROR
	switch char {
		case '+':
			t = scanner.PLUS
		case '-':
			t = scanner.MINUS
		case '*':
			t = scanner.TIMES
		case '(':
			t = scanner.LPAREN
		case ')':
			t = scanner.RPAREN
		default:
			ok = false
	}
	tok = &scanner.Token{t, string(char)}
	return
}
