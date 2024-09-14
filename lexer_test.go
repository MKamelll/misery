package main

import "testing"

func Test_Lexing(t *testing.T) {
	source := "+ - / * let const potato \"string\" 1555 1894.525 class function for if else and or not () [] {}"
	l := NewLexer(source)
	answers := []Token{
		NewToken(TT_Plus, "+"),
		NewToken(TT_Minus, "-"),
		NewToken(TT_Slash, "/"),
		NewToken(TT_Star, "*"),
		NewToken(TT_Let, "let"),
		NewToken(TT_Const, "const"),
		NewToken(TT_Identifier, "potato"),
		NewToken(TT_String, "string"),
		NewToken(TT_Int, "1555"),
		NewToken(TT_Float, "1894.525"),
		NewToken(TT_Class, "class"),
		NewToken(TT_Function, "function"),
		NewToken(TT_For, "for"),
		NewToken(TT_If, "if"),
		NewToken(TT_Else, "else"),
		NewToken(TT_And, "and"),
		NewToken(TT_Or, "or"),
		NewToken(TT_Not, "not"),
		NewToken(TT_LeftParen, "("),
		NewToken(TT_RightParen, ")"),
		NewToken(TT_LeftBrace, "["),
		NewToken(TT_RightBrace, "]"),
		NewToken(TT_LeftBracet, "{"),
		NewToken(TT_RightBracket, "}"),
		NewToken(TT_Eof, "Eof"),
	}

	for _, v := range answers {
		token := l.next()
		if token.kind != v.kind || token.lexeme != v.lexeme {
			t.Fatal("Expected " + v.String() + "instead got " + token.String())
		}

	}
}
