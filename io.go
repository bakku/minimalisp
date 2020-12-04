package tinylisp

import "fmt"

// Println is just the println function
// Usage:
// (println "Hello" "World") => "Hello World"
type Println struct{}

// Arity returns infiniteArity for println.
func (p *Println) Arity() int {
	return infiniteArity
}

// Call implements the println function.
func (p *Println) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	for i, arg := range arguments {
		if i != 0 {
			fmt.Print(" ")
		}

		fmt.Print(arg)
	}

	fmt.Println()

	return nil, nil
}

func (p *Println) String() string {
	return "<println>"
}
