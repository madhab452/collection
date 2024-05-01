package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/madhab452/collection/filter/lexer"
	"github.com/madhab452/collection/filter/token"
)

const PROMPT = ">> "

// this takes the input string and produces a list of tokens
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
		fmt.Println("tokens")
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("\t %+v\n", tok)
		}
	}
}
