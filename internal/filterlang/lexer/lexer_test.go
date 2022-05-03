package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestLex(t *testing.T) {
	lex := lexer.NewLexer(`
	name == "another identifier"
	42 != 13.37
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
	testutils.AssertEqualString(t, lexer.TokenTypeNumber.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "42", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeNotEqual.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeNumber.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "13.37", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeEOF.String(), token.Type.String(), "unexpected token type")
}
