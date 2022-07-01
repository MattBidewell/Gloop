package parser

import (
	"testing"

	"github.com/mattbidewell/gloop-terpreter/ast"
	"github.com/mattbidewell/gloop-terpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 828384;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	if program == nil {
		t.Fatalf("ParserProgram() returned nil")
	}
	if len(program.Statments) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statments))
	}

	tests := []struct {
		expectedIdentifier string
	} {
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statments[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatment. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not s '%s' got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %s. got=%s", name, letStmt.TokenLiteral())
		return false
	}

	return true
}