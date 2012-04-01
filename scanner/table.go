// Package defines a table-driven scanner based on a DFA
package table

import "./scanner"
//import "fmt"

const (
	SPACE = iota
	NEWLINE
	SLASH
	STAR
	LPAREN
	RPAREN
	PLUS
	MINUS
	COLON
	EQUAL
	DOT
	DIGIT
	LETTER
	OTHER
	
	NUM_STATES = 18
)

type row []int
type scanTable []row

func newRow(vals []int) row {
	return row(vals)
}

func newScanTable() *scanTable {
	var s scanTable
	s = append(s, newRow([]int{16, 16, 1, 9, 5, 6, 7, 8, 10, -1, 12, 13, 15, -1}))
	s = append(s, newRow([]int{-1, -1, 2, 3, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{2, 17, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}))
	s = append(s, newRow([]int{3, 3, 3, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}))
	s = append(s, newRow([]int{3, 3, 17, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, 11, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 14, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 14, 13, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 13, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 15, 15, -1}))
	s = append(s, newRow([]int{16, 16, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	s = append(s, newRow([]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	return &s
}

func (s *scanTable) lookup(state, t int) int {
	if state < 0 || state >= NUM_STATES || t < 0 || t > OTHER {
		return -1
	}
	return (*s)[state][t]
}

type Scanner struct {
	input []byte
	pos int
	scanTab *scanTable
	tokenTab map[int]int	// maps states to token-types
	tokens []*scanner.Token
}

func NewScanner(input []byte) *Scanner {
	s := &Scanner{input, 0, newScanTable(), make(map[int]int), []*scanner.Token{}}
	s.tokenTab[0] = scanner.ERROR
	s.tokenTab[1] = scanner.DIV
	s.tokenTab[2] = scanner.ERROR
	s.tokenTab[3] = scanner.ERROR
	s.tokenTab[4] = scanner.ERROR
	s.tokenTab[5] = scanner.LPAREN
	s.tokenTab[6] = scanner.RPAREN
	s.tokenTab[7] = scanner.PLUS
	s.tokenTab[8] = scanner.MINUS
	s.tokenTab[9] = scanner.TIMES
	s.tokenTab[10] = scanner.ERROR
	s.tokenTab[11] = scanner.ASSIGN
	s.tokenTab[12] = scanner.ERROR
	s.tokenTab[13] = scanner.NUMBER
	s.tokenTab[14] = scanner.NUMBER
	s.tokenTab[15] = scanner.ID
	s.tokenTab[16] = scanner.SPACE
	s.tokenTab[17] = scanner.COMMENT
	return s
}

func (s *Scanner) GetToken(i int) *scanner.Token {
	if i > -1 && i < len(s.tokens) {
		return s.tokens[i]
	}
	return nil
}

func (s *Scanner) Scan() (*scanner.Token, bool) {
	tok := s.nextToken()
	for ; tok.Type != scanner.END; tok = s.nextToken() {
		if tok.Type == scanner.ERROR {
			return tok, false
		} else if tok.Type != scanner.SPACE && tok.Type != scanner.COMMENT {
			s.tokens = append(s.tokens, tok)
		}
		if tok.Type == scanner.ID {
			if tok.Value == "read" {
				tok.Type = scanner.READ
			} else if tok.Value == "write" {
				tok.Type = scanner.WRITE
			}
		}
	}
	return nil, true
}

func (s *Scanner) nextToken() *scanner.Token {
	if s.pos >= len(s.input) {
		return &scanner.Token{scanner.END, ""}
	}
	
	state := 0
	value := ""
	char := s.input[s.pos]
	index := s.scanTab.lookup(state, charType(char))
	
	for ; index != -1; index = s.scanTab.lookup(state, charType(char)) {
		state = index
		value += string(char)
		s.pos++
		if s.pos >= len(s.input) {
			break
		}
		char = s.input[s.pos]
	}
	
	if s.tokenTab[state] == scanner.ERROR {
		value = "character '" + string(char) + "' not recognised"
	}
	
	return &scanner.Token{s.tokenTab[state], value}
}

func charType(char byte) int {
	switch char {
		case ' ', '\t':
			return SPACE
		case '\n':
			return NEWLINE
		case '/':
			return SLASH
		case '*':
			return STAR
		case '(':
			return LPAREN
		case ')':
			return RPAREN
		case '+':
			return PLUS
		case '-':
			return MINUS
		case ':':
			return COLON
		case '=':
			return EQUAL
		case '.':
			return DOT
	}
	if scanner.IsDigit(char) {
		return DIGIT
	} else if scanner.IsLetter(char) {
		return LETTER
	}
	return OTHER
}


