package ast_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/ast"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestSimpleCondition(t *testing.T) {
	runner := func(name string, comparisonToken lexer.Token, expectedComparison ast.Comparison) {
		t.Run(name, func(t *testing.T) {
			lex := NewFakeLexer([]lexer.Token{
				lexer.NewTokenIdentifier("http.status"),
				comparisonToken,
				lexer.NewTokenNumber("200"),
			})

			expectedAST := ast.AST{
				Condition: ast.NewCondition(
					ast.NewLabelValue("http.status"),
					expectedComparison,
					ast.NewNumberValue(200),
				),
			}

			actualAST, err := ast.Parse(lex)
			testutils.RequireNoError(t, err, "expecting to parse lexed tokens")

			if expectedAST.Condition != actualAST.Condition {
				t.Errorf("invalid AST\nexpected:\n%v\n\nactual:\n%v\n", expectedAST, actualAST)
			}
		})
	}

	runner("simpleEqualCondition", lexer.NewTokenEqual(), ast.ComparisonEqual)
	runner("simpleNotEqualCondition", lexer.NewTokenNotEqual(), ast.ComparisonNotEqual)
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