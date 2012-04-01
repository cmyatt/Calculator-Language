// Package defines a simple recursive-descent parser
package basic

import (
	"./parser"
	"./../../misc/error"
	"./../scanner/scanner"
)

// Parser is a simple recursive descent parser which does
// not save the resulting parse tree in memory.
type Parser struct {
	sc scanner.Interface
	pos int
	tok *scanner.Token
}

func NewParser(s scanner.Interface) *Parser {
	return &Parser{s, 0, nil}
}

func (p *Parser) getType() int {
	t := scanner.END
	if p.tok != nil {
		t = p.tok.Type
	}
	return t
}

// Updates current token, in the process verifying it is as expected
func (p *Parser) match(nt string, expected int) error.Interface {
	if p.getType() != expected {
		return parser.NewError(nt, []int{expected}, p.tok.Type)
	}
	p.pos++
	p.tok = p.sc.GetToken(p.pos)
	return error.NewGeneric(parser.NO_ERROR, "", "")
}

func (p *Parser) Parse() error.Interface {
	err, ok := p.sc.Scan()
	if !ok {
		return error.NewGeneric("Scanning error", err.Value, "Parser.sc.Scan()")
	}
	p.tok = p.sc.GetToken(p.pos)
	return p.program()
}

func (p *Parser) program() error.Interface {
	switch p.getType() {
	case scanner.ID, scanner.READ, scanner.WRITE, scanner.END:
		err := p.stmtList()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.match("program", scanner.END)
	}
	return parser.NewError("program", []int{scanner.ID, scanner.READ, scanner.WRITE, scanner.END}, p.getType())
}

func (p *Parser) stmtList() error.Interface {
	switch p.getType() {
	case scanner.ID, scanner.READ, scanner.WRITE:
		err := p.stmt()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.stmtList()
	case scanner.END:
		return error.NewGeneric(parser.NO_ERROR, "", "")
	}
	return parser.NewError("stmtList", []int{scanner.ID, scanner.READ, scanner.WRITE, scanner.END}, p.getType())
}

func (p *Parser) stmt() error.Interface {
	switch p.getType() {
	case scanner.ID:
		err := p.match("stmt", scanner.ID)
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		err = p.match("stmt", scanner.ASSIGN)
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.expr()
	case scanner.READ:
		err := p.match("stmt", scanner.READ)
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.match("stmt", scanner.ID)
	case scanner.WRITE:
		err := p.match("stmt", scanner.WRITE)
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.expr()
	}
	return parser.NewError("stmt", []int{scanner.ID, scanner.READ, scanner.WRITE}, p.getType())
}

func (p *Parser) expr() error.Interface {
	switch p.getType() {
	case scanner.ID, scanner.NUMBER, scanner.LPAREN:
		err := p.term()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.termTail()
	}
	return parser.NewError("expr", []int{scanner.ID, scanner.NUMBER, scanner.LPAREN}, p.getType())
}

func (p *Parser) term() error.Interface {
	switch p.getType() {
	case scanner.ID, scanner.NUMBER, scanner.LPAREN:
		err := p.factor()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.factorTail()
	}
	return parser.NewError("term", []int{scanner.ID, scanner.NUMBER, scanner.LPAREN}, p.getType())
}

func (p *Parser) termTail() error.Interface {
	switch p.getType() {
	case scanner.PLUS, scanner.MINUS:
		err := p.addOp()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		err = p.term()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.termTail()
	case scanner.RPAREN, scanner.ID, scanner.READ, scanner.WRITE, scanner.END:
		return error.NewGeneric(parser.NO_ERROR, "", "")
	}
	return parser.NewError("termTail", []int{scanner.PLUS, scanner.MINUS, scanner.RPAREN, scanner.ID,
										   scanner.READ, scanner.WRITE, scanner.END}, p.getType())
}

func (p *Parser) factor() error.Interface {
	switch p.getType() {
	case scanner.ID:
		return p.match("factor", scanner.ID)
	case scanner.NUMBER:
		return p.match("factor", scanner.NUMBER)
	case scanner.LPAREN:
		err := p.match("factor", scanner.LPAREN)
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		err = p.expr()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.match("factor", scanner.RPAREN)
	}
	return parser.NewError("factor", []int{scanner.ID, scanner.NUMBER, scanner.LPAREN}, p.getType())
}

func (p *Parser) factorTail() error.Interface {
	switch p.getType() {
	case scanner.TIMES, scanner.DIV:
		err := p.multOp()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		err = p.factor()
		if err.GetType() != parser.NO_ERROR {
			return err
		}
		return p.factorTail()
	case scanner.PLUS, scanner.MINUS, scanner.RPAREN, scanner.ID, scanner.READ, scanner.WRITE, scanner.END:
		return error.NewGeneric(parser.NO_ERROR, "", "")
	}
	return parser.NewError("factorTail", []int{scanner.TIMES, scanner.DIV, scanner.PLUS, scanner.MINUS,
											 scanner.RPAREN, scanner.ID, scanner.READ, scanner.WRITE,
											 scanner.END}, p.getType())
}

func (p *Parser) addOp() error.Interface {
	switch p.getType() {
	case scanner.PLUS:
		return p.match("addOp", scanner.PLUS)
	case scanner.MINUS:
		return p.match("addOp", scanner.MINUS)
	}
	return parser.NewError("addOp", []int{scanner.PLUS, scanner.MINUS}, p.getType())
}

func (p *Parser) multOp() error.Interface {
	switch p.getType() {
	case scanner.TIMES:
		return p.match("multOp", scanner.TIMES)
	case scanner.DIV:
		return p.match("multOp", scanner.DIV)
	}
	return parser.NewError("multOp", []int{scanner.TIMES, scanner.DIV}, p.getType())
}
