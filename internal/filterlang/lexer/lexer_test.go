package lexer_test

import (
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/filterlang/lexer"
)

func TestLex(t *testing.T) {
	lex := lexer.NewLexer(`
	name == "another identifier"
	several != line
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
	testutils.AssertEqualString(t, lexer.TokenTypeIdentifier.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "several", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeNotEqual.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeIdentifier.String(), token.Type.String(), "unexpected token type")
	testutils.AssertEqualString(t, "line", token.Value, "unexpected token value")

	token = lex.NextToken()
	testutils.AssertEqualString(t, lexer.TokenTypeEOF.String(), token.Type.String(), "unexpected token type")
}
