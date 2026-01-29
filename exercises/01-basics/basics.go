package basics

// Exercise 1: Variables and Types
//
// Coming from JS/TS, practice Go's type system and variable declarations.
// Run tests with: go test -v

// 1. Declare and return a string greeting using shorthand (:=)
// In JS: const greeting = "Hello, Go!"
func GetGreeting() string {
	// TODO: declare greeting using := and return it
	return ""
}

// 2. Return multiple values (name and age)
// In JS: return { name: "Alice", age: 30 } or return ["Alice", 30]
// In Go: functions can return multiple values directly
func GetPersonInfo() (string, int) {
	// TODO: return name "Alice" and age 30
	return "", 0
}

// 3. Type conversion - convert int to float64 percentage
// In JS: const result = num / 100 (automatic)
// In Go: explicit conversion required
func IntToPercentage(n int) float64 {
	// TODO: convert n to float64 and divide by 100
	return 0
}

// 4. Return zero values for each type
// In JS: undefined or null
// In Go: each type has a specific zero value
func GetZeroValues() (int, string, bool, float64) {
	// TODO: declare variables without initializing, return them
	// Hint: var i int (don't assign anything)
	return 0, "", false, 0
}

// 5. Calculate circle area using a constant
// In JS: const PI = 3.14159; return PI * radius * radius
func GetCircleArea(radius float64) float64 {
	// TODO: declare PI as a constant and calculate area
	return 0
}

// 6. Swap two integers and return them
// In JS: return [b, a] or [a, b] = [b, a]
// In Go: multiple return values make this elegant
func Swap(a, b int) (int, int) {
	// TODO: return b, a (swapped)
	return 0, 0
}

// 7. Type inference - Go infers types from values
// Return the type name as a string for learning purposes
func InferredTypes() (intVal int, floatVal float64, stringVal string, boolVal bool) {
	// TODO: use := to declare variables with these values:
	// 42, 3.14, "hello", true
	// Then return them
	return
}
