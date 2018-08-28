package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey_interpreter/lexer"
	"monkey_interpreter/token"
)

// PROMPT represents our prompt in the interpreter
const PROMPT = ">> "

// Start functions starts our interpreter
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
