package tinylisp

import "fmt"

// List is the interface for list implementations to fulfil.
type List interface {
	First() interface{}
	Rest() List
	Add(el interface{}) List
	Len() int
}

// ArrayList represents a simple array list implementation.
type ArrayList struct {
	elements []interface{}
}

// First returns the first element of the ArrayList.
func (a *ArrayList) First() interface{} {
	return a.elements[0]
}

// Rest returns all except the first element of the ArrayList
func (a *ArrayList) Rest() List {
	return NewArrayList(a.elements[1:])
}

// Add returns all except the first element of the ArrayList
func (a *ArrayList) Add(el interface{}) List {
	newElements := a.elements

	newElements = append(newElements, el)

	return NewArrayList(newElements)
}

// Len returns the length of the list.
func (a *ArrayList) Len() int {
	return len(a.elements)
}

// NewArrayList is a factory function to create a new array list.
func NewArrayList(elements []interface{}) *ArrayList {
	return &ArrayList{elements}
}

func (a *ArrayList) String() string {
	var ret string = "("

	for i, el := range a.elements {
		if i != 0 {
			ret += " "
		}

		ret += fmt.Sprintf("%v", el)
	}

	return ret + ")"
}
