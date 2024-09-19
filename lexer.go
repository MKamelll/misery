package main

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	TT_LeftParen        TokenType = "TT_LeftParen"
	TT_RightParen       TokenType = "TT_RightParen"
	TT_LeftBrace        TokenType = "TT_LeftBrace"
	TT_RightBrace       TokenType = "TT_RightBrace"
	TT_LeftBracet       TokenType = "TT_LeftBracet"
	TT_RightBracket     TokenType = "TT_RightBracket"
	TT_Plus             TokenType = "TT_Plus"
	TT_PlusEqual        TokenType = "TT_PlusEqual"
	TT_MinusEqual       TokenType = "TT_MinusEqual"
	TT_Minus            TokenType = "TT_Minus"
	TT_Star             TokenType = "TT_Star"
	TT_StarEqual        TokenType = "TT_StarEqual"
	TT_Slash            TokenType = "TT_Slash"
	TT_SlashEqual       TokenType = "TT_SlashEqual"
	TT_Int              TokenType = "TT_Int"
	TT_Float            TokenType = "TT_Float"
	TT_String           TokenType = "TT_String"
	TT_Identifier       TokenType = "TT_Identifier"
	TT_Eof              TokenType = "TT_Eof"
	TT_Illegal          TokenType = "TT_Illegal"
	TT_GreaterThan      TokenType = "TT_GreaterThan"
	TT_LessThan         TokenType = "TT_LessThan"
	TT_LessThanEqual    TokenType = "TT_LessThanEqual"
	TT_GreaterThanEqual TokenType = "TT_GreaterThanEqual"
	TT_Equal            TokenType = "TT_Equal"
	TT_EqualEqual       TokenType = "TT_EqualEqual"
	TT_Let              TokenType = "TT_Let"
	TT_Const            TokenType = "TT_Const"
	TT_For              TokenType = "TT_For"
	TT_Function         TokenType = "TT_Function"
	TT_Class            TokenType = "TT_Class"
	TT_While            TokenType = "TT_While"
	TT_If               TokenType = "TT_If"
	TT_Else             TokenType = "TT_Else"
	TT_And              TokenType = "TT_And"
	TT_Or               TokenType = "TT_Or"
	TT_Not              TokenType = "TT_Not"
	TT_Carrot           TokenType = "TT_Carrot"
	TT_Colon            TokenType = "TT_Colon"
	TT_Semicolon        TokenType = "TT_Semicolon"
)

type Token struct {
	kind   TokenType
	lexeme string
}

func newToken(kind TokenType, lexeme string) Token {
	return Token{kind: kind, lexeme: lexeme}
}

func (t *Token) String() string {
	return "Token(type: " + string(t.kind) + ", lexeme: '" + t.lexeme + "')"
}

type Lexer struct {
	source     string
	curr_index int
	row        int
	column     int
}

func NewLexer(source string) Lexer {
	return Lexer{source: source, curr_index: 0, row: 0, column: 0}
}

func (l *Lexer) curr() byte {
	return l.source[l.curr_index]
}

func (l *Lexer) advance() {
	l.column++
	l.curr_index++
}

func (l *Lexer) is_at_end() bool {
	return l.curr_index >= len(l.source)
}

func (l *Lexer) next() Token {
	if l.is_at_end() {
		return newToken(TT_Eof, "Eof")
	}

	switch l.curr() {
	case '(':
		{
			l.advance()
			return newToken(TT_LeftParen, "(")
		}
	case ')':
		{
			l.advance()
			return newToken(TT_RightParen, ")")
		}
	case '[':
		{
			l.advance()
			return newToken(TT_LeftBrace, "[")
		}
	case ']':
		{
			l.advance()
			return newToken(TT_RightBrace, "]")
		}
	case '{':
		{
			l.advance()
			return newToken(TT_LeftBracet, "{")
		}
	case '}':
		{
			l.advance()
			return newToken(TT_RightBracket, "}")
		}
	case '+':
		{
			l.advance()
			if !l.is_at_end() && l.curr() == '=' {
				l.advance()
				return newToken(TT_PlusEqual, "+=")
			}

			return newToken(TT_Plus, "+")
		}

	case '-':
		{
			l.advance()
			if !l.is_at_end() && l.curr() == '=' {
				l.advance()
				return newToken(TT_MinusEqual, "-=")
			}

			return newToken(TT_Minus, "-")
		}
	case '*':
		{
			l.advance()
			if !l.is_at_end() && l.curr() == '=' {
				l.advance()
				return newToken(TT_StarEqual, "*=")
			}

			return newToken(TT_Star, "*")
		}
	case '/':
		{
			l.advance()
			if !l.is_at_end() && l.curr() == '=' {
				l.advance()
				return newToken(TT_SlashEqual, "/=")
			}

			return newToken(TT_Slash, "/")
		}
	case '>':
		{
			l.advance()
			if !l.is_at_end() && l.curr() == '=' {
				l.advance()
				return newToken(TT_GreaterThanEqual, ">=")
			}

			return newToken(TT_GreaterThan, ">")
		}
	case '<':
		{
			l.advance()
			if !l.is_at_end() && l.curr() == '=' {
				l.advance()
				return newToken(TT_LessThanEqual, "<=")
			}

			return newToken(TT_LessThan, "<")
		}
	case '=':
		{
			l.advance()
			if !l.is_at_end() && l.curr() == '=' {
				l.advance()
				return newToken(TT_EqualEqual, "==")
			}

			return newToken(TT_Equal, "=")
		}
	case ' ':
		{
			l.advance()
			return l.next()
		}
	case '\n':
		{
			l.advance()
			l.column = 0
			l.row++
			return l.next()
		}

	case '"':
		{
			result := l.is_string()
			return newToken(TT_String, result)
		}
	case ':':
		{
			l.advance()
			return newToken(TT_Colon, ":")
		}
	case ';':
		{
			l.advance()
			return newToken(TT_Semicolon, ";")
		}
	case '^':
		{
			l.advance()
			return newToken(TT_Carrot, "^")
		}
	default:
		{
			if unicode.IsDigit(rune(l.curr())) {
				number, is_float := l.is_number()
				number = strings.Trim(number, "\n")
				if is_float {
					return newToken(TT_Float, number)
				}
				return newToken(TT_Int, number)
			}

			if unicode.IsLetter(rune(l.curr())) {
				identifier := l.is_identifier()
				switch identifier {
				case "if":
					{
						return newToken(TT_If, identifier)
					}
				case "for":
					{
						return newToken(TT_For, identifier)
					}
				case "while":
					{
						return newToken(TT_While, identifier)
					}
				case "function":
					{
						return newToken(TT_Function, identifier)
					}
				case "class":
					{
						return newToken(TT_Class, identifier)
					}
				case "and":
					{
						return newToken(TT_And, identifier)
					}
				case "or":
					{
						return newToken(TT_Or, identifier)
					}
				case "not":
					{
						return newToken(TT_Not, identifier)
					}
				case "let":
					{
						return newToken(TT_Let, identifier)
					}
				case "const":
					{
						return newToken(TT_Const, identifier)
					}
				case "else":
					{
						return newToken(TT_Else, identifier)
					}
				default:
					break
				}

				return newToken(TT_Identifier, identifier)
			}

			break
		}
	}

	illegal := l.curr()
	l.advance()
	return newToken(TT_Illegal, string(illegal))
}

func (l *Lexer) is_number() (string, bool) {
	result := ""
	is_bool := false

	for !l.is_at_end() {
		if l.curr() == '.' {
			is_bool = true
			result += string(l.curr())
			l.advance()
			continue
		}

		if !unicode.IsDigit(rune(l.curr())) {
			break
		}

		result += string(l.curr())
		l.advance()

	}

	return result, is_bool

}

func (l *Lexer) is_string() string {
	l.advance()
	result := ""

	for !l.is_at_end() {
		if l.curr() == '"' {
			l.advance()
			break
		}

		result += string(l.curr())
		l.advance()
	}

	return result
}

func (l *Lexer) is_identifier() string {
	result := ""

	for !l.is_at_end() {
		if !unicode.IsDigit(rune(l.curr())) && !unicode.IsLetter(rune(l.curr())) {
			break
		}

		result += string(l.curr())
		l.advance()
	}

	return result
}
