package main

import (
	"testing"
)

func Test_Parsing(t *testing.T) {

	source := "potato \"string\" 1555 1894.525"
	l := NewLexer(source)
	p := NewParser(l)
	trees, err := p.Parse()
	if err != nil {
		panic(err)
	}

	answers := []Expression{
		NewIdentifierExpr("potato"),
		NewStringExpr("string"),
		NewIntExpr(1555),
		NewFloatExpr(1894.525),
	}

	cond1 := trees[0].(IdentifierExpr) == answers[0].(IdentifierExpr)
	cond2 := trees[1].(StringExpr) == answers[1].(StringExpr)
	cond3 := trees[2].(IntExpr) == answers[2].(IntExpr)
	cond4 := trees[3].(FloatExpr) == answers[3].(FloatExpr)

	if !cond1 {
		t.Fatal(cond1)
	} else if !cond2 {
		t.Fatal(cond2)
	} else if !cond3 {
		t.Fatal(cond3)
	} else if !cond4 {
		t.Fatal(cond4)
	}
}
