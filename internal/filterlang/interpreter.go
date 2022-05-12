package filterlang

import (
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/lonepeon/tailog/internal/filterlang/ast"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

type Interpreter struct {
	tree ast.AST
}

func Parse(source string) (Interpreter, error) {
	tree, err := ast.Parse(lexer.NewLexer(source))
	if err != nil {
		return Interpreter{}, err
	}

	return Interpreter{tree: tree}, nil
}

func (r Interpreter) Execute(entry decoding.Entry) bool {
	return execute(entry, r.tree.Condition)
}

func execute(entry decoding.Entry, condition ast.Condition) bool {
	switch cond := condition.(type) {
	case ast.ConditionAnd:
		return execute(entry, cond.Left()) && execute(entry, cond.Right())
	case ast.ConditionOr:
		return execute(entry, cond.Left()) || execute(entry, cond.Right())
	case ast.ConditionExpression:
		return executeExpression(entry, cond)
	}

	return false
}

func getField(entry decoding.Entry, valuer ast.Valuer) (decoding.Field, bool) {
	switch v := valuer.(type) {
	case ast.FieldValue:
		return entry.Field(v.Value())
	case ast.StringValue:
		return decoding.NewFieldString("dummy", v.Value()), true
	case ast.NumberValue:
		return decoding.NewFieldNumber("dummy", v.Value()), true
	default:
		return nil, false
	}
}

func executeExpression(entry decoding.Entry, cond ast.ConditionExpression) bool {
	field1, ok := getField(entry, cond.Left())
	if !ok {
		return false
	}

	field2, ok := getField(entry, cond.Right())
	if !ok {
		return false
	}

	comp := field1.Compare(field2)

	if comp == decoding.FieldComparisonEqual && cond.Comparison() == ast.ComparisonEqual {
		return true
	}

	if comp != decoding.FieldComparisonEqual && cond.Comparison() == ast.ComparisonNotEqual {
		return true
	}

	return false
}
