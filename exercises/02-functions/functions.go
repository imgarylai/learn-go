package functions

// Exercise 2: Functions and Error Handling
//
// Practice Go's function syntax and explicit error handling.
// No try/catch here - errors are values!
// Run tests with: go test -v

import "errors"

// 1. Multiple return values - return quotient and remainder
// In JS: return { quotient, remainder } or return [quotient, remainder]
func Divide(a, b int) (int, int) {
	// TODO: return a/b and a%b
	return 0, 0
}

// 2. Named return values with naked return
// Go lets you name return values and use "return" without arguments
func DivideNamed(a, b int) (quotient, remainder int) {
	// TODO: assign to quotient and remainder, then just "return"
	return
}

// 3. Error handling - the Go way
// In JS: throw new Error("cannot divide by zero")
// In Go: return error as a value
func SafeDivide(a, b int) (int, error) {
	// TODO: if b is 0, return 0 and errors.New("cannot divide by zero")
	// Otherwise return a/b and nil
	return 0, nil
}

// 4. Functions as values (first-class functions)
// In JS: const add = (a, b) => a + b
func GetOperation(op string) func(int, int) int {
	// TODO: return a function based on op:
	// "add" -> returns a + b
	// "subtract" -> returns a - b
	// "multiply" -> returns a * b
	// default -> returns function that returns 0
	return nil
}

// 5. Variadic functions (like JS rest parameters)
// In JS: function sum(...numbers) { return numbers.reduce((a,b) => a+b, 0) }
func Sum(numbers ...int) int {
	// TODO: sum all numbers using range
	return 0
}

// 6. Closure - function that captures outer variable
// In JS: const counter = () => { let count = 0; return () => ++count; }
func MakeCounter() func() int {
	// TODO: return a function that increments and returns a counter
	// Each call should return 1, 2, 3, ...
	return nil
}

// 7. Higher-order function - takes a function as parameter
// In JS: array.map(fn)
func MapInts(numbers []int, fn func(int) int) []int {
	// TODO: apply fn to each number and return new slice
	return nil
}

// Keep import used
var _ = errors.New
