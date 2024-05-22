// Package compiler provides functionality for compiling AST nodes into bytecode instructions.
//
// It emits the result of the compilation, including the emitted instructions and the constant pool.
package compiler

import (
	"chui/ast"
	"chui/code"
	"chui/object"
	"fmt"
)

// Compiler is the compiler struct - it holds the bytecode instructions and the constant pool.
type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

// New() creates a new Compiler.
func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

// Compile() compiles an AST node into bytecode instructions.
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {

	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)

			if err != nil {
				return err
			}
		}

	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)

		if err != nil {
			return err
		}

		c.emit(code.OpPop)

	case *ast.InfixExpression:
		err := c.Compile(node.Left)

		if err != nil {
			return err
		}

		err = c.Compile(node.Right)

		if err != nil {
			return err
		}

		switch node.Operator {
		case "+":
			c.emit(code.OpAdd)

		case "-":
			c.emit(code.OpSub)

		case "*":
			c.emit(code.OpMul)

		case "/":
			c.emit(code.OpDiv)

		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}

	case *ast.Boolean:
		if node.Value {
			c.emit(code.OpTrue)
		} else {
			c.emit(code.OpFalse)
		}

	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))
	}

	return nil
}

// addConstant() adds a constant to the constant pool and returns its index.
func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)

	return len(c.constants) - 1
}

// emit() appends an instruction to the bytecode.
func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)

	return pos
}

// addInstruction() adds an instruction to the bytecode.
func (c *Compiler) addInstruction(ins []byte) int {
	posNewInstruction := len(c.instructions)
	c.instructions = append(c.instructions, ins...)

	return posNewInstruction
}

// Bytecode() returns the compiled bytecode.
func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// Bytecode is the Bytecode struct - it holds the bytecode instructions and the constant pool.
type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
