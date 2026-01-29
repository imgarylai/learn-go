package interfaces

// Exercise 5: Interfaces
//
// Go interfaces are implicit - no "implements" keyword!
// If a type has the right methods, it implements the interface.
// Think duck typing with compile-time safety.
// Run tests with: go test -v

import (
	"fmt"
	"math"
)

// Shape interface - any type with Area() and Perimeter() is a Shape
// In TS: interface Shape { Area(): number; Perimeter(): number; }
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape (implicitly!)
type Rectangle struct {
	Width  float64
	Height float64
}

// 1. Implement Area for Rectangle
func (r Rectangle) Area() float64 {
	// TODO: return width * height
	return 0
}

// 2. Implement Perimeter for Rectangle
func (r Rectangle) Perimeter() float64 {
	// TODO: return 2 * (width + height)
	return 0
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

// 3. Implement Area for Circle (use math.Pi)
func (c Circle) Area() float64 {
	// TODO: return Pi * radius^2
	return 0
}

// 4. Implement Perimeter for Circle (circumference)
func (c Circle) Perimeter() float64 {
	// TODO: return 2 * Pi * radius
	return 0
}

// 5. Function that works with ANY Shape
// This is the power of interfaces!
func DescribeShape(s Shape) string {
	// TODO: return "Area: X.XX, Perimeter: X.XX"
	// Use fmt.Sprintf with %.2f format
	return ""
}

// 6. Type assertion - check if interface is specific type
// In TS: value as Type or <Type>value
func GetRadius(s Shape) (float64, bool) {
	// TODO: if s is a Circle, return its radius and true
	// Otherwise return 0 and false
	// Hint: circle, ok := s.(Circle)
	return 0, false
}

// 7. Type switch - handle different types
func DescribeType(s Shape) string {
	// TODO: return "Rectangle" if Rectangle, "Circle" if Circle, "Unknown" otherwise
	// Use type switch: switch v := s.(type) { case Rectangle: ... }
	return ""
}

// Stringer interface - like toString() in JS
// fmt package uses this when printing
type Person struct {
	Name string
	Age  int
}

// 8. Implement Stringer for Person
// Return format: "Name (Age years old)"
func (p Person) String() string {
	// TODO: return formatted string
	return ""
}

// error interface - Go's way of handling errors
// Just needs Error() string method
type ValidationError struct {
	Field   string
	Message string
}

// 9. Implement error interface for ValidationError
// Return format: "validation failed on FIELD: MESSAGE"
func (e ValidationError) Error() string {
	// TODO: return formatted error message
	return ""
}

// 10. Function that returns our custom error
func ValidateName(name string) error {
	// TODO: if name is empty, return ValidationError{Field: "name", Message: "required"}
	// Otherwise return nil
	return nil
}

// Empty interface (any) - accepts any type
// In TS: any or unknown

// 11. Type assertion with any
func StringLength(v any) int {
	// TODO: if v is a string, return its length
	// Otherwise return -1
	return -1
}

// 12. Handle multiple types with type switch
func Describe(v any) string {
	// TODO: return description based on type:
	// int: "integer: X"
	// string: "string: X"
	// bool: "boolean: X"
	// default: "unknown"
	return ""
}

// Keep imports used
var _ = math.Pi
var _ = fmt.Sprintf
