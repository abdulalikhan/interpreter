package lexer

import (
	"github.com/abdulalikhan/interpreter/token"
)

type Lexer struct {
	input     string
	currIndex int
	currChar  byte // a byte is being used for this field - this limits the chars in our language to ASCII
	nextIndex int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}

	// when creating a new lexer, call readChar() to initialize the indices and currChar
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextIndex >= len(l.input) {
		l.currChar = 0
	} else {
		l.currChar = l.input[l.nextIndex]
	}

	// update the lexer's current index and increment the next index
	l.currIndex = l.nextIndex
	l.nextIndex++
}

func newToken(tokenType token.TokenType, value byte) token.Token {
	return token.Token{Type: tokenType, Value: string(value)}
}

// isCharAllowedInIdentifier dictates if a character is allowed as part of an identifier
// letters a-z, A-Z and the _ character are allowed by the programming language
func isCharAllowedInIdentifier(char byte) bool {
	if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char == '_' {
		return true
	}
	return false
}

func isDigit(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

// consumeWhitespace ignores whitespace characters in the lexer's feed
func (l *Lexer) consumeWhitespace() {
	for l.currChar == ' ' || l.currChar == '\n' || l.currChar == '\t' || l.currChar == '\r' {
		l.readChar()
	}
}

// readIndentifier reads an identifier until a non-letter character is encountered
func (l *Lexer) readIdentifier() string {
	startIndex := l.currIndex

	for isCharAllowedInIdentifier(l.currChar) {
		l.readChar()
	}

	return l.input[startIndex:l.currIndex]
}

// readNumber reads a number until a non-numeric character is encountered
func (l *Lexer) readNumber() string {
	startIndex := l.currIndex

	for isDigit(l.currChar) {
		l.readChar()
	}

	return l.input[startIndex:l.currIndex]
}

func (l *Lexer) NextToken() token.Token {
	var nextToken token.Token

	l.consumeWhitespace()

	switch l.currChar {
	case '=':
		nextToken = newToken(token.ASSIGN, l.currChar)
	case ';':
		nextToken = newToken(token.SEMICOLON, l.currChar)
	case '(':
		nextToken = newToken(token.LPAREN, l.currChar)
	case ')':
		nextToken = newToken(token.RPAREN, l.currChar)
	case '{':
		nextToken = newToken(token.LBRACE, l.currChar)
	case '}':
		nextToken = newToken(token.RBRACE, l.currChar)
	case '+':
		nextToken = newToken(token.PLUS, l.currChar)
	case ',':
		nextToken = newToken(token.COMMA, l.currChar)
	case 0:
		nextToken.Type = token.EOF
		nextToken.Value = ""
	default:
		// if currChar is a valid identifier character - the lexer has detected an identifier
		if isCharAllowedInIdentifier(l.currChar) {
			nextToken.Value = l.readIdentifier()
			nextToken.Type = token.GetIdentifierTypeByValue(nextToken.Value)
			// readIdentifier sets the currIndex ahead of the identifier - do not call readChar at the end
			return nextToken
		} else if isDigit(l.currChar) {
			nextToken.Value = l.readNumber()
			nextToken.Type = token.INT
			// readNumber sets the currIndex ahead of the integer - do not call readChar at the end
			return nextToken
		} else {
			nextToken = newToken(token.ILLEGAL, l.currChar)
		}
	}

	l.readChar()
	return nextToken
}
