package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

// PROMPT is printed at the beginning of every line
const PROMPT = ">>"

// Start is the main loop to run the repl
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			// Without this fmt.Println(), the terminal prompt stays on the same
			// line as the repl prompt when exiting using ctrl+d or ctrl+c.
			fmt.Println()
			return
		}

		line := scanner.Text()

		if line == "exit" {
			return
		}

		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}
}
