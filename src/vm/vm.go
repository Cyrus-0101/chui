// Package vm provides functionality for executing bytecode instructions.
//
// It includes functionality for creating a new VM, executing bytecode, and manipulating the VM's stack.
package vm

import (
	"chui/code"
	"chui/compiler"
	"chui/object"
	"fmt"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions
	stack        []object.Object
	sp           int // Always points to the next value. Top of stack is stack[sp-1]
}

// New() creates a new VM with the given bytecode.
func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        make([]object.Object, StackSize), // The stack is preallocated to have a StackSize number of elements, which should be enough for us
		sp:           0,                                // sp will always point to the next free slot in the stack. If there’s one element on the stack, located at index 0, the value of sp would be 1 and to access the element we’d use stack[sp-1].
	}
}

// StackTop() pushes an object onto the stack.
func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}

	return vm.stack[vm.sp-1]
}

// Run() executes the VM's bytecode instructions. It iterates over the instructions, decodes the current instruction, and executes the corresponding operation.
func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		switch op {

		case code.OpConstant:
			// code.ReadUint16(): I am fast :D
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			err := vm.push(vm.constants[constIndex])

			if err != nil {
				return err
			}

		case code.OpAdd:
			right := vm.pop()
			left := vm.pop()

			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			result := leftValue + rightValue
			vm.push(&object.Integer{Value: result})

		}
	}
	return nil
}

// push() checks the stack size pushes an object onto the stack.
func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}

// pop() pops an object off the stack, from the top .
func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--

	return o
}
