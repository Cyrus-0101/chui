// Package code provides functionality for working with bytecode instructions.
//
// It defines the Instructions type, which is a slice of bytes, and the Opcode type, which is a byte.
package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
	OpAdd
	OpPop
	OpSub
	OpMul
	OpDiv
)

// Definition represents the definition of an opcode, including its name and the widths of its operands, which is used to determine how many bytes to read to extract the operands.
type Definition struct {
	Name          string
	OperandWidths []int
}

// definitions maps opcodes to their definitions.
var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	OpAdd:      {"OpAdd", []int{}},
	OpPop:      {"OpPop", []int{}},
	OpSub:      {"OpSub", []int{}},
	OpMul:      {"OpMul", []int{}},
	OpDiv:      {"OpDiv", []int{}},
}

// Lookup() retrieves the definition of an opcode.
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

// Make() creates a bytecode instruction from an opcode and its operands.
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]

	if !ok {
		return []byte{}
	}

	instructionLen := 1

	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)
	offset := 1

	for i, o := range operands {
		width := def.OperandWidths[i]

		switch width {

		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}

		offset += width
	}

	return instruction
}

// String() returns a string representation of the bytecode instructions, including the offset of each instruction in the bytecode.
func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0

	for i < len(ins) {
		def, err := Lookup(ins[i])

		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))

		i += 1 + read
	}

	return out.String()
}

// fmtInstruction() formats an instruction for printing.
func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}

	switch operandCount {

	case 0:
		return def.Name

	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

// ReadOperands() reads the operands of an instruction.
func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {

		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

// ReadUint16() reads a uint16 from a byte slice.
func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
