package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/santos-404/myte/lexer"
	"github.com/santos-404/myte/parser"
)


const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {  // This is a common while true loop
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}
		
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, err := range errors {
		io.WriteString(out, "\t"+err+"\n")
	}
}
