package lexer

import (
	"testing"

	"github.com/abdulalikhan/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `let number_a = 2;
			let number_b = 3;
			let sum = fn(a, b) {
			  a + b;
			};
			let result = sum(number_a, number_b);
      if (result == 5) {
        return true;
      } else {
        return false;
      }`

	testCases := []struct {
		expectedTokenType token.TokenType
		expectedValue     string
	}{
		{token.LET, "let"},
		{token.IDENT, "number_a"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "number_b"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "sum"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "sum"},
		{token.LPAREN, "("},
		{token.IDENT, "number_a"},
		{token.COMMA, ","},
		{token.IDENT, "number_b"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.EQ, "=="},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	lex := NewLexer(input)
	for i, testCase := range testCases {
		thisToken := lex.NextToken()
		if thisToken.Type != testCase.expectedTokenType {
			t.Fatalf("test #%d failed - incorrect token type. expected=%q, got=%q", i, testCase.expectedTokenType, thisToken.Type)
		}

		if thisToken.Value != testCase.expectedValue {
			t.Fatalf("test #%d failed - incorrect token value. expected=%q, got=%q", i, testCase.expectedValue, thisToken.Value)
		}
	}
}
