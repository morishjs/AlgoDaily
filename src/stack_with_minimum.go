package main

import (
	"errors"
)

// Recall that a stack is an abstract data type modeling a collection elements. Its primary operations are push (which adds an element to the top of the stack) and pop (which removes the most newest element).

// Traditionally, a stack can easily be implemented in Javascript and many other languages using an array (and its built-in methods).

// const stack = [];

// stack.push(5);

// stack.push(6);

// stack.push(7);

// stack.pop();

// // 7

// stack.pop();

// // 6

// However, let's say we wanted to implement a stack with the following interface:

// push(val) - add an element on to the top of the stack.

// pop(val) - remove the element at the top of the stack and return it.

// peek(val) - see the element at the top of the stack without removing it.

// min() - get minimum element in stack.

// How would you do it, and can you implement it via a MinStack class? The class should have the following methods. Work off this skeleton:

// class MinStack {

//   constructor() {

//   }

//   push(val) {

//   }

//   pop() {

//   }

//   peek() {

//   }

//   min() {

//   }

// }

// Can you do call min() and retrieve it in O(1) time?

// solution: https://algodaily.com/challenges/implement-a-stack-with-minimum

type MinStack struct {
	minStack []int
	minValues []int
}

func NewMinStack() *MinStack {
	stack := new(MinStack)

	return stack
}

func (s *MinStack) push(val int) {
	s.minStack = append(s.minStack, val)

	minValue := s.minValues[len(s.minValues) - 1]

	if len(s.minStack) == 1 || minValue > val {
		s.minValues = append(s.minValues, val)
	} else {
		s.minValues = append(s.minValues, minValue)
	}
}

func (s *MinStack) pop() (int, error) {
	if s.empty() == true {
		return 0, errors.New("Stack is empty")
	}

	stackLength := len(s.minStack)
	minValuesLength := len(s.minValues)

	poppedValue := s.minStack[stackLength-1]
	
	s.minStack = s.minStack[:stackLength-1]
	s.minValues = s.minValues[:minValuesLength-1]
	
	return poppedValue, nil
}

func (s *MinStack) peek() (int, error) {
	if s.empty() == true {
		return 0, errors.New("Stack is empty")
	}

	stackLength := len(s.minStack)
	topValue := s.minStack[stackLength-1]

	return topValue, nil
}

func (s *MinStack) empty() bool {
	stackLength := len(s.minStack)

	if stackLength == 0 {
		return true
	}

	return false
}

func (s *MinStack) min() (int, error) {
	if s.empty() == true {
		return 0, errors.New("Stack is empty")
	}

	return s.minValues[len(s.minValues)-1], nil
}
