// Package provides structs, interfaces and constants for use in different types of scanners
package scanner

const (
	ASSIGN = iota
	PLUS
	MINUS
	TIMES
	DIV
	LPAREN
	RPAREN
	ID
	READ
	WRITE
	NUMBER
	SPACE
	COMMENT
	ERROR
	END
	NUM_TOKENS
)

type Token struct {
	Type int
	Value string
}

type Interface interface {
	Scan() (*Token, bool)		// if error, return it and false, else return nil and true
	GetToken(i int) *Token
}

func IsDigit(char byte) bool {
	switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return true
	}
	return false
}

func IsLetter(char byte) bool {
	switch char {
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			return true
	}
	return false
}

func GetTokenName(tok int) string {
	switch tok {
		case ASSIGN:
			return "Assign"
		case PLUS:
			return "Plus"
		case MINUS:
			return "Minus"
		case TIMES:
			return "Times"
		case DIV:
			return "Div"
		case LPAREN:
			return "Lparen"
		case RPAREN:
			return "Rparen"
		case ID:
			return "Id"
		case READ:
			return "Read"
		case WRITE:
			return "Write"
		case NUMBER:
			return "Number"
		case ERROR:
			return "Error"
		case END:
			return "End"
	}
	return "undefined token"
}
