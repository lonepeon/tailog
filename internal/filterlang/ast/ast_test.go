package ast_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/ast"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestCondition(t *testing.T) {
	type TestCase struct {
		Tokens      []lexer.Token
		ExpectedAST ast.AST
	}

	runner := func(name string, tc TestCase) {
		t.Run(name, func(t *testing.T) {
			lex := NewFakeLexer(tc.Tokens)
			actualAST, err := ast.Parse(lex)
			testutils.RequireNoError(t, err, "expecting to parse lexed tokens")

			if tc.ExpectedAST.Condition != actualAST.Condition {
				t.Errorf("invalid AST\nexpected:\n%v\n\nactual:\n%v\n", tc.ExpectedAST, actualAST)
			}
		})
	}

	runner("simpleEqualCondition", TestCase{
		Tokens: []lexer.Token{
			lexer.NewTokenIdentifier("http.status"),
			lexer.NewTokenEqual(),
			lexer.NewTokenNumber("200"),
		},
		ExpectedAST: ast.AST{
			Condition: ast.NewCondition(
				ast.NewLabelValue("http.status"),
				ast.ComparisonEqual,
				ast.NewNumberValue(200),
			),
		},
	})
}

type Lexer struct {
	tokens []lexer.Token
	index  int
}

func NewFakeLexer(tokens []lexer.Token) *Lexer {
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
