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
				lexer.NewTokenField("http.status"),
				comparisonToken,
				lexer.NewTokenNumber("200"),
			})

			expectedAST := ast.AST{
				Condition: ast.NewConditionExpression(
					ast.NewFieldValue("http.status"),
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

func TestAndOperator(t *testing.T) {
	lex := NewFakeLexer([]lexer.Token{
		lexer.NewTokenField("http.status"),
		lexer.NewTokenEqual(),
		lexer.NewTokenNumber("200"),
		lexer.NewTokenAnd(),
		lexer.NewTokenField("user.id"),
		lexer.NewTokenNotEqual(),
		lexer.NewTokenNumber("42"),
	})

	expectedAST := ast.AST{
		Condition: ast.NewConditionAnd(
			ast.NewConditionExpression(
				ast.NewFieldValue("http.status"),
				ast.ComparisonEqual,
				ast.NewNumberValue(200),
			),
			ast.NewConditionExpression(
				ast.NewFieldValue("user.id"),
				ast.ComparisonNotEqual,
				ast.NewNumberValue(42),
			),
		),
	}

	actualAST, err := ast.Parse(lex)
	testutils.RequireNoError(t, err, "expecting to parse lexed tokens")

	if expectedAST.Condition != actualAST.Condition {
		t.Errorf("invalid AST\nexpected:\n%v\n\nactual:\n%v\n", expectedAST, actualAST)
	}
}

func TestOrOperator(t *testing.T) {
	lex := NewFakeLexer([]lexer.Token{
		lexer.NewTokenField("http.status"),
		lexer.NewTokenEqual(),
		lexer.NewTokenNumber("200"),
		lexer.NewTokenOr(),
		lexer.NewTokenField("user.id"),
		lexer.NewTokenNotEqual(),
		lexer.NewTokenNumber("42"),
	})

	expectedAST := ast.AST{
		Condition: ast.NewConditionOr(
			ast.NewConditionExpression(
				ast.NewFieldValue("http.status"),
				ast.ComparisonEqual,
				ast.NewNumberValue(200),
			),
			ast.NewConditionExpression(
				ast.NewFieldValue("user.id"),
				ast.ComparisonNotEqual,
				ast.NewNumberValue(42),
			),
		),
	}

	actualAST, err := ast.Parse(lex)
	testutils.RequireNoError(t, err, "expecting to parse lexed tokens")

	if expectedAST.Condition != actualAST.Condition {
		t.Errorf("invalid AST\nexpected:\n%v\n\nactual:\n%v\n", expectedAST, actualAST)
	}
}

func TestAndOrOperators(t *testing.T) {
	lex := NewFakeLexer([]lexer.Token{
		lexer.NewTokenField("http.status"),
		lexer.NewTokenEqual(),
		lexer.NewTokenNumber("200"),
		lexer.NewTokenAnd(),
		lexer.NewTokenField("http.method"),
		lexer.NewTokenNotEqual(),
		lexer.NewTokenString("POST"),
		lexer.NewTokenOr(),
		lexer.NewTokenField("user.id"),
		lexer.NewTokenNotEqual(),
		lexer.NewTokenNumber("1337"),
	})

	expectedAST := ast.AST{
		Condition: ast.NewConditionOr(
			ast.NewConditionAnd(
				ast.NewConditionExpression(
					ast.NewFieldValue("http.status"),
					ast.ComparisonEqual,
					ast.NewNumberValue(200),
				),
				ast.NewConditionExpression(
					ast.NewFieldValue("http.method"),
					ast.ComparisonNotEqual,
					ast.NewStringValue("POST"),
				),
			),
			ast.NewConditionExpression(
				ast.NewFieldValue("user.id"),
				ast.ComparisonNotEqual,
				ast.NewNumberValue(1337),
			),
		),
	}

	actualAST, err := ast.Parse(lex)
	testutils.RequireNoError(t, err, "expecting to parse lexed tokens")

	if expectedAST.Condition != actualAST.Condition {
		t.Errorf("invalid AST\nexpected:\n%v\n\nactual:\n%v\n", expectedAST, actualAST)
	}
}

func TestOrAndOperators(t *testing.T) {
	lex := NewFakeLexer([]lexer.Token{
		lexer.NewTokenField("user.id"),
		lexer.NewTokenNotEqual(),
		lexer.NewTokenNumber("42"),
		lexer.NewTokenOr(),
		lexer.NewTokenField("http.status"),
		lexer.NewTokenEqual(),
		lexer.NewTokenNumber("200"),
		lexer.NewTokenAnd(),
		lexer.NewTokenField("http.method"),
		lexer.NewTokenNotEqual(),
		lexer.NewTokenString("POST"),
	})

	expectedAST := ast.AST{
		Condition: ast.NewConditionOr(
			ast.NewConditionExpression(
				ast.NewFieldValue("user.id"),
				ast.ComparisonNotEqual,
				ast.NewNumberValue(42),
			),
			ast.NewConditionAnd(
				ast.NewConditionExpression(
					ast.NewFieldValue("http.status"),
					ast.ComparisonEqual,
					ast.NewNumberValue(200),
				),
				ast.NewConditionExpression(
					ast.NewFieldValue("http.method"),
					ast.ComparisonNotEqual,
					ast.NewStringValue("POST"),
				),
			),
		),
	}

	actualAST, err := ast.Parse(lex)
	testutils.RequireNoError(t, err, "expecting to parse lexed tokens")

	if expectedAST.Condition != actualAST.Condition {
		t.Errorf("invalid AST\nexpected:\n%v\n\nactual:\n%v\n", expectedAST, actualAST)
	}
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
