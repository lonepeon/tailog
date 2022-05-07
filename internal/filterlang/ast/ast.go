package ast

import (
	"fmt"
	"strconv"

	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

type Lexer interface {
	NextToken() lexer.Token
}

type Comparison struct {
	name string
}

func (s Comparison) String() string {
	return s.name
}

var (
	ComparisonUndefined = Comparison{name: "undefined"}
	ComparisonEqual     = Comparison{name: "equal"}
)

type LabelValue struct {
	Value string
}

func (l LabelValue) String() string {
	return fmt.Sprintf("label(%q)", l.Value)
}

func (l LabelValue) isValue() {}

type NumberValue struct {
	Value float64
}

func (n NumberValue) String() string {
	return fmt.Sprintf("%f", n.Value)
}

func (n NumberValue) isValue() {}

type Valuer interface {
	isValue()
	String() string
}

type Condition struct {
	Left       Valuer
	Comparison Comparison
	Right      Valuer
}

func (c Condition) String() string {
	return fmt.Sprintf("%s %s %s", c.Left, c.Comparison, c.Right)
}

type AST struct {
	Condition Condition
}

func (a AST) String() string {
	return a.Condition.String()
}

func From(lex Lexer) (AST, error) {
	left, err := readValue(lex.NextToken())
	if err != nil {
		return AST{}, fmt.Errorf("can't parse left side of the condition: %w", err)
	}

	comparison, err := readComparison(lex.NextToken())
	if err != nil {
		return AST{}, fmt.Errorf("can't parse condition comparison: %w", err)
	}

	right, err := readValue(lex.NextToken())
	if err != nil {
		return AST{}, fmt.Errorf("can't parse right side of the condition: %w", err)
	}

	return AST{
		Condition: Condition{
			Left:       left,
			Comparison: comparison,
			Right:      right,
		},
	}, nil
}

func readValue(token lexer.Token) (Valuer, error) {
	switch token.Type {
	case lexer.TokenTypeIdentifier:
		return LabelValue{Value: token.Value}, nil
	case lexer.TokenTypeNumber:
		number, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse %q to number: %v", token.Value, err)
		}
		return NumberValue{Value: number}, nil
	default:
		return nil, fmt.Errorf("expecting an Identifier or Number type of token but got %s", token)
	}
}

func readComparison(token lexer.Token) (Comparison, error) {
	switch token.Type {
	case lexer.TokenTypeEqual:
		return ComparisonEqual, nil
	default:
		return ComparisonUndefined, fmt.Errorf("expecting an Equal token but got %s", token)
	}
}
