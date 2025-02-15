package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/eval"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">>"

const MONKEY_FACE = `
_______AAAA___r_o_h_i_t___AAAA________
       VVVV               VVVV        
       (__)               (__)
        \ \               / /
         \ \   \\|||//   / /
          > \   _   _   / <
 hang      > \ / \ / \ / <
  in        > \\_o_o_// <
  there...   > ( (_) ) <
              >|     |<
             / |\___/| \
             / (_____) \
             /         \
              /   o   \
               ) ___ (   
              / /   \ \  
             ( /     \ )
             ><       ><
            ///\     /\\\
            '''       '''
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan() // read input from user
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := eval.Eval(program)

		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
