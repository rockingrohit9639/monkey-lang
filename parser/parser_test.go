package parser_test

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"testing"
)

func TestLestStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not return. got=%q", returnStmt.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not let, got=%q", s.TokenLiteral())
		return false
	}

	letSmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%q", s)
		return false
	}

	if letSmt.Name.Value != name {
		t.Errorf("letSmt.Name.Value not %s. got=%q", name, letSmt.Name.Value)
		return false
	}

	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%q", name, letSmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
