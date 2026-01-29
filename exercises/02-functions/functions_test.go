package functions

import (
	"reflect"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     int
		wantQ    int
		wantR    int
	}{
		{17, 5, 3, 2},
		{10, 3, 3, 1},
		{20, 4, 5, 0},
		{7, 2, 3, 1},
	}

	for _, tc := range tests {
		q, r := Divide(tc.a, tc.b)
		if q != tc.wantQ || r != tc.wantR {
			t.Errorf("Divide(%d, %d): got (%d, %d), want (%d, %d)",
				tc.a, tc.b, q, r, tc.wantQ, tc.wantR)
		}
	}
}

func TestDivideNamed(t *testing.T) {
	q, r := DivideNamed(17, 5)
	if q != 3 || r != 2 {
		t.Errorf("DivideNamed(17, 5): got (%d, %d), want (3, 2)", q, r)
	}
}

func TestSafeDivide(t *testing.T) {
	// Test normal division
	result, err := SafeDivide(10, 2)
	if err != nil {
		t.Errorf("SafeDivide(10, 2): unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("SafeDivide(10, 2): got %d, want 5", result)
	}

	// Test division by zero
	_, err = SafeDivide(10, 0)
	if err == nil {
		t.Error("SafeDivide(10, 0): expected error, got nil")
	}
}

func TestGetOperation(t *testing.T) {
	add := GetOperation("add")
	if add == nil {
		t.Fatal("GetOperation(\"add\") returned nil")
	}
	if add(3, 4) != 7 {
		t.Errorf("add(3, 4): got %d, want 7", add(3, 4))
	}

	subtract := GetOperation("subtract")
	if subtract == nil {
		t.Fatal("GetOperation(\"subtract\") returned nil")
	}
	if subtract(10, 4) != 6 {
		t.Errorf("subtract(10, 4): got %d, want 6", subtract(10, 4))
	}

	multiply := GetOperation("multiply")
	if multiply == nil {
		t.Fatal("GetOperation(\"multiply\") returned nil")
	}
	if multiply(3, 4) != 12 {
		t.Errorf("multiply(3, 4): got %d, want 12", multiply(3, 4))
	}

	unknown := GetOperation("unknown")
	if unknown == nil {
		t.Fatal("GetOperation(\"unknown\") returned nil")
	}
	if unknown(3, 4) != 0 {
		t.Errorf("unknown(3, 4): got %d, want 0", unknown(3, 4))
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{10, 20, 30}, 60},
		{[]int{}, 0},
		{[]int{42}, 42},
	}

	for _, tc := range tests {
		result := Sum(tc.input...)
		if result != tc.expected {
			t.Errorf("Sum(%v): got %d, want %d", tc.input, result, tc.expected)
		}
	}
}

func TestMakeCounter(t *testing.T) {
	counter := MakeCounter()
	if counter == nil {
		t.Fatal("MakeCounter() returned nil")
	}

	// Each call should increment
	for i := 1; i <= 5; i++ {
		result := counter()
		if result != i {
			t.Errorf("counter() call %d: got %d, want %d", i, result, i)
		}
	}

	// New counter should start fresh
	counter2 := MakeCounter()
	if counter2() != 1 {
		t.Error("new counter should start at 1")
	}
}

func TestMapInts(t *testing.T) {
	// Test double
	result := MapInts([]int{1, 2, 3}, func(n int) int { return n * 2 })
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MapInts double: got %v, want %v", result, expected)
	}

	// Test square
	result = MapInts([]int{1, 2, 3, 4}, func(n int) int { return n * n })
	expected = []int{1, 4, 9, 16}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MapInts square: got %v, want %v", result, expected)
	}

	// Test empty
	result = MapInts([]int{}, func(n int) int { return n })
	if len(result) != 0 {
		t.Errorf("MapInts empty: got %v, want empty slice", result)
	}
}
