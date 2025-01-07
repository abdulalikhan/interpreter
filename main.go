package main

import (
	"fmt"
	"os"

	"github.com/abdulalikhan/interpreter/repl"
)

func main() {
	fmt.Printf("Welcome to the fictitious programming language interpreter\n")
	repl.Start(os.Stdin, os.Stdout)
}
