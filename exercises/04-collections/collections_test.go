package collections

import (
	"reflect"
	"sort"
	"testing"
)

func TestCreateSlice(t *testing.T) {
	result := CreateSlice()
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestSliceMiddle(t *testing.T) {
	result := SliceMiddle([]int{1, 2, 3, 4, 5})
	expected := []int{2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}

	// Test short slice
	result = SliceMiddle([]int{1, 2})
	if len(result) != 0 {
		t.Errorf("short slice: got %v, want empty", result)
	}
}

func TestDouble(t *testing.T) {
	result := Double([]int{1, 2, 3})
	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}

	// Test empty
	result = Double([]int{})
	if len(result) != 0 {
		t.Errorf("empty: got %v, want empty", result)
	}
}

func TestFilterGreaterThan(t *testing.T) {
	result := FilterGreaterThan([]int{1, 5, 10, 3, 8, 2}, 5)
	expected := []int{10, 8}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0},
		{[]int{100}, 100},
		{[]int{-1, 1}, 0},
	}

	for _, tc := range tests {
		result := Sum(tc.input)
		if result != tc.expected {
			t.Errorf("Sum(%v): got %d, want %d", tc.input, result, tc.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 5, 3, 9, 2}, 9},
		{[]int{-5, -1, -10}, -1},
		{[]int{42}, 42},
		{[]int{}, 0},
	}

	for _, tc := range tests {
		result := Max(tc.input)
		if result != tc.expected {
			t.Errorf("Max(%v): got %d, want %d", tc.input, result, tc.expected)
		}
	}
}

func TestCreateScores(t *testing.T) {
	scores := CreateScores()

	if scores == nil {
		t.Fatal("CreateScores returned nil")
	}

	expected := map[string]int{"alice": 95, "bob": 87, "charlie": 92}
	if !reflect.DeepEqual(scores, expected) {
		t.Errorf("got %v, want %v", scores, expected)
	}
}

func TestGetScore(t *testing.T) {
	scores := map[string]int{"alice": 95, "bob": 87}

	score, exists := GetScore(scores, "alice")
	if !exists || score != 95 {
		t.Errorf("GetScore(alice): got (%d, %v), want (95, true)", score, exists)
	}

	score, exists = GetScore(scores, "unknown")
	if exists {
		t.Errorf("GetScore(unknown): got exists=true, want false")
	}
}

func TestGetTopScorer(t *testing.T) {
	scores := map[string]int{"alice": 95, "bob": 87, "charlie": 92}
	result := GetTopScorer(scores)

	if result != "alice" {
		t.Errorf("got %q, want %q", result, "alice")
	}

	// Test empty map
	result = GetTopScorer(map[string]int{})
	if result != "" {
		t.Errorf("empty map: got %q, want empty string", result)
	}
}

func TestRemovePlayer(t *testing.T) {
	scores := map[string]int{"alice": 95, "bob": 87, "charlie": 92}
	RemovePlayer(scores, "bob")

	if _, exists := scores["bob"]; exists {
		t.Error("bob should have been removed")
	}

	if len(scores) != 2 {
		t.Errorf("expected 2 remaining, got %d", len(scores))
	}
}

func TestCountOccurrences(t *testing.T) {
	items := []string{"a", "b", "a", "c", "b", "a"}
	result := CountOccurrences(items)

	expected := map[string]int{"a": 3, "b": 2, "c": 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestGetAdults(t *testing.T) {
	people := []Person{
		{"Alice", 25},
		{"Bob", 17},
		{"Charlie", 30},
		{"Diana", 15},
	}

	result := GetAdults(people)

	if len(result) != 2 {
		t.Errorf("expected 2 adults, got %d", len(result))
	}

	for _, p := range result {
		if p.Age < 18 {
			t.Errorf("found minor in adults: %+v", p)
		}
	}
}

func TestGetNames(t *testing.T) {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}

	result := GetNames(people)
	expected := []string{"Alice", "Bob"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestFindByName(t *testing.T) {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}

	result := FindByName(people, "Alice")
	if result == nil {
		t.Fatal("FindByName(Alice) returned nil")
	}
	if result.Name != "Alice" || result.Age != 25 {
		t.Errorf("got %+v, want Alice/25", result)
	}

	result = FindByName(people, "Unknown")
	if result != nil {
		t.Errorf("FindByName(Unknown): got %+v, want nil", result)
	}
}

// Keep import used
var _ = sort.Strings
