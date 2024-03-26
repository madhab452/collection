package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/madhab452/collection/filter/lexer"
	"github.com/madhab452/collection/filter/parser"
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

		fmt.Println("statements")
		l2 := lexer.New(line)
		p := parser.New(l2)
		f := p.ParseFilter()
		if errors := p.Errors(); len(errors) > 0 {
			for i := 0; i < len(errors); i++ {
				fmt.Printf("\t error %d, %s \n", i+1, errors[i])
			}
		} else {
			for _, stmt := range f.Statements {
				fmt.Printf("\t [field: %s , Operator: %s , Value: %s (%s)] \n", stmt.Field, stmt.Operator, stmt.Value.Literal, stmt.Value.ValueType)
			}
		}

	}
}
