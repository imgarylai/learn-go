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
	return a / b, a % b
}

// 2. Named return values with naked return
// Go lets you name return values and use "return" without arguments
func DivideNamed(a, b int) (quotient, remainder int) {
	// TODO: assign to quotient and remainder, then just "return"
	quotient = a / b
	remainder = a % b
	return
}

// 3. Error handling - the Go way
// In JS: throw new Error("cannot divide by zero")
// In Go: return error as a value
func SafeDivide(a, b int) (int, error) {
	// TODO: if b is 0, return 0 and errors.New("cannot divide by zero")
	// Otherwise return a/b and nil
	if b != 0 {
		return a / b, nil
	}
	return 0, errors.New("cannot divide by zero")
}

// 4. Functions as values (first-class functions)
// In JS: const add = (a, b) => a + b
func GetOperation(op string) func(int, int) int {
	// TODO: return a function based on op:
	// "add" -> returns a + b
	// "subtract" -> returns a - b
	// "multiply" -> returns a * b
	// default -> returns function that returns 0
	switch op {
	case "add":
		return func(a int, b int) int { return a + b }
	case "subtract":
		return func(a int, b int) int { return a - b }
	case "multiply":
		return func(a int, b int) int { return a * b }
	}
	return func(a int, b int) int { return 0 }
}

// 5. Variadic functions (like JS rest parameters)
// In JS: function sum(...numbers) { return numbers.reduce((a,b) => a+b, 0) }
func Sum(numbers ...int) int {
	// TODO: sum all numbers using range
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// 6. Closure - function that captures outer variable
// In JS: const counter = () => { let count = 0; return () => ++count; }
func MakeCounter() func() int {
	// TODO: return a function that increments and returns a counter
	// Each call should return 1, 2, 3, ...
	n := 0
	return func() int { n++; return n }
}

// 7. Higher-order function - takes a function as parameter
// In JS: array.map(fn)
func MapInts(numbers []int, fn func(int) int) []int {
	// TODO: apply fn to each number and return new slice
	// res := []int{}
	// for _, v := range numbers {
	// 	res = append(res, fn(v))
	// }
	// return res
	res := make([]int, len(numbers)) // Pre-allocate exact size
	for i, v := range numbers {
		res[i] = fn(v) // Direct assignment, no append
	}
	return res
}

// Keep import used
var _ = errors.New
