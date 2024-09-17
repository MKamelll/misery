package main

import (
	"fmt"
	"strconv"
)

type Expression interface {
	String() string
}

type IntExpr struct {
	value int
}

func NewIntExpr(value int) IntExpr {
	return IntExpr{value: value}
}

func (i IntExpr) String() string {
	return fmt.Sprintf("IntExpr(value: %d)", i.value)
}

type FloatExpr struct {
	value float64
}

func NewFloatExpr(value float64) FloatExpr {
	return FloatExpr{value: value}
}

func (f FloatExpr) String() string {
	return fmt.Sprintf("FloatExpr(value: %f)", f.value)
}

type StringExpr struct {
	value string
}

func NewStringExpr(value string) StringExpr {
	return StringExpr{value: value}
}

func (s StringExpr) String() string {
	return fmt.Sprintf("StringExpr(value: '%s')", s.value)
}

type IdentifierExpr struct {
	value string
}

func NewIdentifierExpr(value string) IdentifierExpr {
	return IdentifierExpr{value: value}
}

func (i IdentifierExpr) String() string {
	return fmt.Sprintf("IdentifierExpr(value: '%s')", i.value)
}

type BinaryExpr struct {
	lhs Expression
	op  string
	rhs Expression
}

func NewBinaryExpr(lhs Expression, op string, rhs Expression) BinaryExpr {
	return BinaryExpr{lhs: lhs, op: op, rhs: rhs}
}

func (b BinaryExpr) String() string {
	return fmt.Sprintf("BinaryExpr(lhs: %s, op: %s, rhs: %s)", b.lhs.String(), b.op, b.rhs.String())
}

type Associativity = string

const (
	Associativity_Left  Associativity = "Associativity_Left"
	Associativity_Right Associativity = "Associativity_Right"
)

type Operator struct {
	op            string
	precedence    int
	associativity Associativity
}

func newOperator(op string, precedence int, associativity Associativity) Operator {
	return Operator{op: op, precedence: precedence, associativity: associativity}
}

var allowed_ops = map[string]Operator{
	"+": newOperator("+", 1, Associativity_Left),
	"-": newOperator("-", 1, Associativity_Left),
	"*": newOperator("*", 2, Associativity_Left),
	"/": newOperator("/", 2, Associativity_Left),
}

type Parser struct {
	lexer        Lexer
	curr_token   Token
	prev_token   Token
	parsed_trees []Expression
}

func NewParser(lexer Lexer) Parser {
	return Parser{lexer: lexer, curr_token: lexer.next()}
}

func (p *Parser) advance() {
	p.prev_token = p.curr_token
	p.curr_token = p.lexer.next()
}

func (p *Parser) is_at_end() bool {
	return p.curr_token.kind == TT_Eof
}

func (p *Parser) match(kinds ...TokenType) bool {
	for _, kind := range kinds {
		if p.curr_token.kind == kind {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) Parse() ([]Expression, error) {
	if p.is_at_end() {
		return p.parsed_trees, nil
	}

	expr, err := p.parse_expr(0)
	if err != nil {
		return p.parsed_trees, err
	}
	p.parsed_trees = append(p.parsed_trees, expr)

	return p.Parse()
}

func (p *Parser) parse_expr(min_precedence int) (Expression, error) {
	lhs, err := p.parse_int_expr()
	if err != nil {
		return nil, err
	}

	for !p.is_at_end() {
		op := p.curr_token.lexeme
		val, ok := allowed_ops[op]
		if !ok || val.precedence < min_precedence {
			break
		}

		next_min_precedence := val.precedence

		if val.associativity == Associativity_Left {
			next_min_precedence += 1
		}

		p.advance()
		rhs, err := p.parse_expr(next_min_precedence)

		if err != nil {
			return nil, err
		}

		lhs = NewBinaryExpr(lhs, op, rhs)
	}

	return lhs, nil

}

func (p *Parser) parse_int_expr() (Expression, error) {
	if p.match(TT_Int) {
		num, err := strconv.ParseInt(p.prev_token.lexeme, 10, 64)
		if err != nil {
			return nil, err
		}
		return NewIntExpr(int(num)), nil
	}

	return p.parse_float_expr()
}

func (p *Parser) parse_float_expr() (Expression, error) {
	if p.match(TT_Float) {
		num, err := strconv.ParseFloat(p.prev_token.lexeme, 64)

		if err != nil {
			return nil, err
		}

		return NewFloatExpr(num), nil
	}

	return p.parse_string_expr()
}

func (p *Parser) parse_string_expr() (Expression, error) {
	if p.match(TT_String) {
		return NewStringExpr(p.prev_token.lexeme), nil
	}

	return p.parse_identifier_expr()
}

func (p *Parser) parse_identifier_expr() (Expression, error) {
	if p.match(TT_Identifier) {
		return NewIdentifierExpr(p.prev_token.lexeme), nil
	}

	return nil, fmt.Errorf("%d:%d: unexpected token '%s'", p.lexer.row, p.lexer.column, p.curr_token.lexeme)
}
