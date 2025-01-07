package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/abdulalikhan/interpreter/lexer"
	"github.com/abdulalikhan/interpreter/token"
)

const PROMPT = ">> "

func Start(inStream io.Reader, outStream io.Writer) {
	scanner := bufio.NewScanner(inStream)
	for {
		fmt.Printf(PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)
		for thisToken := l.NextToken(); thisToken.Type != token.EOF; thisToken = l.NextToken() {
			fmt.Printf("%+v\n", thisToken)
		}
	}
}
