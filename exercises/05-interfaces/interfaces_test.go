package interfaces

import (
	"math"
	"strings"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()

	if area != 50 {
		t.Errorf("got %f, want 50", area)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	perimeter := rect.Perimeter()

	if perimeter != 30 {
		t.Errorf("got %f, want 30", perimeter)
	}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 5}
	area := circle.Area()

	expected := math.Pi * 25 // Pi * r^2
	if math.Abs(area-expected) > 0.0001 {
		t.Errorf("got %f, want %f", area, expected)
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := Circle{Radius: 5}
	perimeter := circle.Perimeter()

	expected := 2 * math.Pi * 5 // 2 * Pi * r
	if math.Abs(perimeter-expected) > 0.0001 {
		t.Errorf("got %f, want %f", perimeter, expected)
	}
}

func TestShapeInterface(t *testing.T) {
	// Both Rectangle and Circle should implement Shape
	var s Shape

	s = Rectangle{Width: 10, Height: 5}
	if s.Area() != 50 {
		t.Error("Rectangle should implement Shape")
	}

	s = Circle{Radius: 5}
	if s.Area() == 0 {
		t.Error("Circle should implement Shape")
	}
}

func TestDescribeShape(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	desc := DescribeShape(rect)

	if !strings.Contains(desc, "50.00") {
		t.Errorf("should contain area 50.00, got %q", desc)
	}
	if !strings.Contains(desc, "30.00") {
		t.Errorf("should contain perimeter 30.00, got %q", desc)
	}
}

func TestGetRadius(t *testing.T) {
	circle := Circle{Radius: 7}
	radius, ok := GetRadius(circle)

	if !ok {
		t.Error("should return true for Circle")
	}
	if radius != 7 {
		t.Errorf("got radius %f, want 7", radius)
	}

	rect := Rectangle{Width: 10, Height: 5}
	_, ok = GetRadius(rect)
	if ok {
		t.Error("should return false for Rectangle")
	}
}

func TestDescribeType(t *testing.T) {
	tests := []struct {
		shape    Shape
		expected string
	}{
		{Rectangle{}, "Rectangle"},
		{Circle{}, "Circle"},
	}

	for _, tc := range tests {
		result := DescribeType(tc.shape)
		if result != tc.expected {
			t.Errorf("DescribeType(%T): got %q, want %q", tc.shape, result, tc.expected)
		}
	}
}

func TestPersonString(t *testing.T) {
	person := Person{Name: "Alice", Age: 30}
	str := person.String()

	if !strings.Contains(str, "Alice") {
		t.Errorf("should contain name, got %q", str)
	}
	if !strings.Contains(str, "30") {
		t.Errorf("should contain age, got %q", str)
	}
}

func TestValidationError(t *testing.T) {
	err := ValidationError{Field: "email", Message: "invalid format"}
	errStr := err.Error()

	if !strings.Contains(errStr, "email") {
		t.Errorf("should contain field name, got %q", errStr)
	}
	if !strings.Contains(errStr, "invalid format") {
		t.Errorf("should contain message, got %q", errStr)
	}
}

func TestValidateName(t *testing.T) {
	// Valid name
	err := ValidateName("Alice")
	if err != nil {
		t.Errorf("valid name should return nil, got %v", err)
	}

	// Empty name
	err = ValidateName("")
	if err == nil {
		t.Error("empty name should return error")
	}

	// Check it's our custom error type
	if _, ok := err.(ValidationError); !ok {
		t.Error("should return ValidationError type")
	}
}

func TestStringLength(t *testing.T) {
	tests := []struct {
		input    any
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{42, -1},
		{true, -1},
	}

	for _, tc := range tests {
		result := StringLength(tc.input)
		if result != tc.expected {
			t.Errorf("StringLength(%v): got %d, want %d", tc.input, result, tc.expected)
		}
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		input    any
		contains string
	}{
		{42, "integer"},
		{"hello", "string"},
		{true, "boolean"},
		{3.14, "unknown"},
	}

	for _, tc := range tests {
		result := Describe(tc.input)
		if !strings.Contains(result, tc.contains) {
			t.Errorf("Describe(%v): got %q, should contain %q", tc.input, result, tc.contains)
		}
	}
}
