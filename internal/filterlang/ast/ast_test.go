package ast_test

import (
	"fmt"
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/ast"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestCondition(t *testing.T) {
	lex := NewFakeLexer(
		lexer.NewTokenIdentifier("http.status"),
		lexer.NewTokenEqual(),
		lexer.NewTokenNumber("200"),
	)

	expectedAST := ast.AST{
		Condition: ast.Condition{
			Left:       ast.LabelValue{Value: "http.status"},
			Comparison: ast.ComparisonEqual,
			Right:      ast.NumberValue{Value: 200},
		},
	}

	actualAST, err := ast.From(lex)
	testutils.RequireNoError(t, err, "expecting to parse lexed tokens")

	assertAST(t, expectedAST, actualAST, "unexpected AST")
}

type Lexer struct {
	tokens []lexer.Token
	index  int
}

func NewFakeLexer(tokens ...lexer.Token) *Lexer {
	return &Lexer{tokens: tokens}
}

func (l *Lexer) NextToken() lexer.Token {
	if l.index >= len(l.tokens) {
		return lexer.NewTokenEOF()
	}

	token := l.tokens[l.index]
	l.index += 1

	return token
}

func assertAST(t *testing.T, expected ast.AST, actual ast.AST, pattern string, vars ...interface{}) {
	if expected.Condition != actual.Condition {
		t.Errorf(fmt.Sprintf("invalid AST: %v\nexpected:\n%v\n\nactual:\n%v\n", fmt.Sprintf(pattern, vars...), expected, actual))
	}
}
