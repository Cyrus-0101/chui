package vm

import (
	"chui/ast"
	"chui/compiler"
	"chui/lexer"
	"chui/object"
	"chui/parser"
	"fmt"
	"testing"
)

// TestIntegerArithmetic() tests whether the VM can perform integer arithmetic.
func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"3 + 2", 5},
		{"50 - 25", 25},
	}

	runVmTests(t, tests)
}

// parse() parses the input string and returns the AST.
func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

// testIntegerObject() tests the integer object.
func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)

	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)",
			actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
	}

	return nil
}

type vmTestCase struct {
	input    string
	expected interface{}
}

// runVmTests() runs the tests for the VM.
func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	// Iterate over the tests.
	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()
		err := comp.Compile(program)

		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()

		if err != nil {
			t.Fatalf("vm error: %s", err)
		}

		stackElem := vm.LastPoppedStackElem()
		testExpectedObject(t, tt.expected, stackElem)

	}
}

// testExpectedObject() checks the expected object against the actual object.
func testExpectedObject(
	t *testing.T,
	expected interface{},
	actual object.Object,
) {
	t.Helper()

	switch expected := expected.(type) {

	// If the expected value is an integer, test the integer object.
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	}
}
