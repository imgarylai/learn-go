package basics

import (
	"math"
	"testing"
)

func TestGetGreeting(t *testing.T) {
	greeting := GetGreeting()

	if greeting == "" {
		t.Error("GetGreeting returned empty string, expected a greeting")
	}

	if greeting != "Hello, Go!" {
		t.Errorf("got %q, want %q", greeting, "Hello, Go!")
	}
}

func TestGetPersonInfo(t *testing.T) {
	name, age := GetPersonInfo()

	if name != "Alice" {
		t.Errorf("name: got %q, want %q", name, "Alice")
	}

	if age != 30 {
		t.Errorf("age: got %d, want %d", age, 30)
	}
}

func TestIntToPercentage(t *testing.T) {
	tests := []struct {
		input    int
		expected float64
	}{
		{42, 0.42},
		{100, 1.0},
		{0, 0.0},
		{50, 0.5},
	}

	for _, tc := range tests {
		result := IntToPercentage(tc.input)
		if result != tc.expected {
			t.Errorf("IntToPercentage(%d): got %f, want %f", tc.input, result, tc.expected)
		}
	}
}

func TestGetZeroValues(t *testing.T) {
	i, s, b, f := GetZeroValues()

	if i != 0 {
		t.Errorf("int zero value: got %d, want 0", i)
	}
	if s != "" {
		t.Errorf("string zero value: got %q, want empty string", s)
	}
	if b != false {
		t.Errorf("bool zero value: got %v, want false", b)
	}
	if f != 0.0 {
		t.Errorf("float64 zero value: got %f, want 0.0", f)
	}
}

func TestGetCircleArea(t *testing.T) {
	tests := []struct {
		radius   float64
		expected float64
	}{
		{1.0, 3.14159},
		{5.0, 78.53975},
		{0.0, 0.0},
	}

	for _, tc := range tests {
		result := GetCircleArea(tc.radius)
		// Allow small floating point difference
		if math.Abs(result-tc.expected) > 0.0001 {
			t.Errorf("GetCircleArea(%f): got %f, want %f", tc.radius, result, tc.expected)
		}
	}
}

func TestSwap(t *testing.T) {
	a, b := Swap(1, 2)

	if a != 2 || b != 1 {
		t.Errorf("Swap(1, 2): got (%d, %d), want (2, 1)", a, b)
	}

	x, y := Swap(100, -50)
	if x != -50 || y != 100 {
		t.Errorf("Swap(100, -50): got (%d, %d), want (-50, 100)", x, y)
	}
}

func TestInferredTypes(t *testing.T) {
	intVal, floatVal, stringVal, boolVal := InferredTypes()

	if intVal != 42 {
		t.Errorf("intVal: got %d, want 42", intVal)
	}
	if floatVal != 3.14 {
		t.Errorf("floatVal: got %f, want 3.14", floatVal)
	}
	if stringVal != "hello" {
		t.Errorf("stringVal: got %q, want %q", stringVal, "hello")
	}
	if boolVal != true {
		t.Errorf("boolVal: got %v, want true", boolVal)
	}
}
