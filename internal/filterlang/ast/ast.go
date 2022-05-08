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
	ComparisonNotEqual  = Comparison{name: "not equal"}
)

type LabelValue struct {
	value string
}

func NewLabelValue(labelName string) LabelValue {
	return LabelValue{value: labelName}
}

func (l LabelValue) Value() string {
	return l.value
}

func (l LabelValue) String() string {
	return fmt.Sprintf("label(%q)", l.value)
}

func (l LabelValue) isValue() {}

type NumberValue struct {
	value float64
}

func NewNumberValue(value float64) NumberValue {
	return NumberValue{value: value}
}

func (l NumberValue) Value() float64 {
	return l.value
}

func (n NumberValue) String() string {
	return fmt.Sprintf("%f", n.value)
}

func (n NumberValue) isValue() {}

type Valuer interface {
	isValue()
	String() string
}

type Condition struct {
	left       Valuer
	comparison Comparison
	right      Valuer
}

func NewCondition(left Valuer, comparison Comparison, right Valuer) Condition {
	return Condition{left: left, comparison: comparison, right: right}
}

func (c Condition) Left() Valuer {
	return c.left
}

func (c Condition) Right() Valuer {
	return c.right
}

func (c Condition) Comparison() Comparison {
	return c.comparison
}

func (c Condition) String() string {
	return fmt.Sprintf("%s %s %s", c.left, c.comparison, c.right)
}

type AST struct {
	Condition Condition
}

func (a AST) String() string {
	return a.Condition.String()
}

func Parse(lex Lexer) (AST, error) {
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

	return AST{Condition: NewCondition(left, comparison, right)}, nil
}

func readValue(token lexer.Token) (Valuer, error) {
	switch token.Type {
	case lexer.TokenTypeIdentifier:
		return NewLabelValue(token.Value), nil
	case lexer.TokenTypeNumber:
		number, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse %q to number: %v", token.Value, err)
		}
		return NewNumberValue(number), nil
	default:
		return nil, fmt.Errorf("expecting an Identifier or Number type of token but got %s", token)
	}
}

func readComparison(token lexer.Token) (Comparison, error) {
	switch token.Type {
	case lexer.TokenTypeEqual:
		return ComparisonEqual, nil
	case lexer.TokenTypeNotEqual:
		return ComparisonNotEqual, nil
	default:
		return ComparisonUndefined, fmt.Errorf("expecting an Equal or NotEqual token but got %s", token)
	}
}
