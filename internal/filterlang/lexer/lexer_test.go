package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

// nolint:funlen
func TestLex(t *testing.T) {
	lex := lexer.NewLexer(`
	lbl:name == lbl:"another identifier" || 12 == lbl:something
	42 != 13.37 && lbl:id == lbl:"something else"
	`)

	token := lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeIdentifier.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "name", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeEqual.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeIdentifier.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "another identifier", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeOr.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeNumber.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "12", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeEqual.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeIdentifier.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "something", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeNumber.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "42", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeNotEqual.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeNumber.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "13.37", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeAnd.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeIdentifier.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "id", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeEqual.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeIdentifier.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "something else", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeEOF.String(), token.Type.String(), "unexpected token type")
}
