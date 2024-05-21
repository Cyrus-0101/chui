// Package object provides functionality for creating and manipulating environments.
//
// It includes functionality for creating new environments, setting values in environments, and retrieving values from environments.
package object

// NewEnclosedEnvironment creates a new enclosed environment with the given outer environment.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment creates a new environment.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment represents an environment with a store of objects and an outer environment.
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get retrieves the value associated with the given name from the environment.
// It returns the value and a boolean indicating whether the value was found.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets the value of the given name in the environment and returns the value.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
