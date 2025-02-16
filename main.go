package main

import (
	"fmt"
	"monkey/eval"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)

	monkeySrcPath, err := filepath.Abs("./main.monkey")
	check(err)

	monkeySrc, err := os.ReadFile(monkeySrcPath)
	check(err)

	monkeySrcStr := string(monkeySrc)

	l := lexer.New(monkeySrcStr)
	p := parser.New(l)
	env := object.NewEnvironment()

	program := p.ParseProgram()

	evaluated := eval.Eval(program, env)

	fmt.Print(evaluated.Inspect())
	fmt.Print("\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
