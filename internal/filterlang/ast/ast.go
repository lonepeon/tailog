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
	ComparisonEqual     = Comparison{name: "=="}
	ComparisonNotEqual  = Comparison{name: "!="}
)

type FieldValue struct {
	value string
}

func NewFieldValue(fieldName string) FieldValue {
	return FieldValue{value: fieldName}
}

func (l FieldValue) Value() string {
	return l.value
}

func (l FieldValue) String() string {
	return fmt.Sprintf("field:%q", l.value)
}

func (l FieldValue) isValue() {}

type NumberValue struct {
	value float64
}

func NewNumberValue(value float64) NumberValue {
	return NumberValue{value: value}
}

func (n NumberValue) Value() float64 {
	return n.value
}

func (n NumberValue) String() string {
	return fmt.Sprintf("%f", n.value)
}

func (n NumberValue) isValue() {}

type StringValue struct {
	value string
}

func NewStringValue(value string) StringValue {
	return StringValue{value: value}
}

func (s StringValue) Value() string {
	return s.value
}

func (s StringValue) String() string {
	return fmt.Sprintf("%q", s.value)
}

func (s StringValue) isValue() {}

type Valuer interface {
	isValue()
	String() string
}

type Condition interface {
	isCondition()
	String() string
}

type ConditionAnd struct {
	left  Condition
	right Condition
}

func NewConditionAnd(left Condition, right Condition) ConditionAnd {
	return ConditionAnd{left: left, right: right}
}

func (c ConditionAnd) isCondition() {}

func (c ConditionAnd) Left() Condition {
	return c.left
}

func (c ConditionAnd) Right() Condition {
	return c.right
}

func (c ConditionAnd) String() string {
	return fmt.Sprintf("(%s && %s)", c.left, c.right)
}

type ConditionOr struct {
	left  Condition
	right Condition
}

func NewConditionOr(left Condition, right Condition) ConditionOr {
	return ConditionOr{left: left, right: right}
}

func (c ConditionOr) isCondition() {}

func (c ConditionOr) Left() Condition {
	return c.left
}

func (c ConditionOr) Right() Condition {
	return c.right
}

func (c ConditionOr) String() string {
	return fmt.Sprintf("(%s || %s)", c.left, c.right)
}

type ConditionExpression struct {
	left       Valuer
	comparison Comparison
	right      Valuer
}

func NewConditionExpression(left Valuer, comparison Comparison, right Valuer) ConditionExpression {
	return ConditionExpression{left: left, comparison: comparison, right: right}
}

func (c ConditionExpression) isCondition() {}

func (c ConditionExpression) Left() Valuer {
	return c.left
}

func (c ConditionExpression) Right() Valuer {
	return c.right
}

func (c ConditionExpression) Comparison() Comparison {
	return c.comparison
}

func (c ConditionExpression) String() string {
	return fmt.Sprintf("%s %s %s", c.left, c.comparison, c.right)
}

type AST struct {
	Condition Condition
}

func (a AST) String() string {
	return a.Condition.String()
}

func Parse(lex Lexer) (AST, error) {
	condition, err := readCondition(lex)
	if err != nil {
		return AST{}, err
	}

	return AST{Condition: condition}, nil
}

func readCondition(lex Lexer) (Condition, error) {
	var condition Condition

	condition, err := readExpression(lex)
	if err != nil {
		return nil, err
	}

	for {
		token := lex.NextToken()
		if token == lexer.NewTokenEOF() {
			return condition, nil
		}

		condition, err = readLogicalOperator(lex, condition, token)
		if err != nil {
			return nil, err
		}
	}
}

func readLogicalOperator(lex Lexer, currentCondition Condition, currentToken lexer.Token) (Condition, error) {
	switch currentToken.Type {
	case lexer.TokenTypeAnd:
		otherExpression, err := readExpression(lex)
		if err != nil {
			return nil, err
		}

		return NewConditionAnd(currentCondition, otherExpression), nil
	case lexer.TokenTypeOr:
		otherExpression, err := readCondition(lex)
		if err != nil {
			return nil, err
		}

		return NewConditionOr(currentCondition, otherExpression), nil
	}

	return nil, fmt.Errorf("expecting a And or Or token but got %s", currentToken)
}

func readExpression(lex Lexer) (ConditionExpression, error) {
	left, err := readValue(lex.NextToken())
	if err != nil {
		return ConditionExpression{}, fmt.Errorf("can't parse left side of the condition: %w", err)
	}

	comparison, err := readComparison(lex.NextToken())
	if err != nil {
		return ConditionExpression{}, fmt.Errorf("can't parse condition comparison: %w", err)
	}

	right, err := readValue(lex.NextToken())
	if err != nil {
		return ConditionExpression{}, fmt.Errorf("can't parse right side of the condition: %w", err)
	}

	return NewConditionExpression(left, comparison, right), nil
}

func readValue(token lexer.Token) (Valuer, error) {
	switch token.Type {
	case lexer.TokenTypeField:
		return NewFieldValue(token.Value), nil
	case lexer.TokenTypeString:
		return NewStringValue(token.Value), nil
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
