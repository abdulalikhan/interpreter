package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

// keywordMap maps literals to language keywords
var keywordMap = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// GetIdentifierTypeByValue translates identifier values to token types
func GetIdentifierTypeByValue(identifierVal string) TokenType {
	if tokenType, ok := keywordMap[identifierVal]; ok {
		return tokenType
	}
	return IDENT
}
