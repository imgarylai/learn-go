package concurrency

import (
	"sort"
	"testing"
	"time"
)

func TestChannelBasics(t *testing.T) {
	result := ChannelBasics()
	if result != 42 {
		t.Errorf("got %d, want 42", result)
	}
}

func TestBufferedChannel(t *testing.T) {
	result := BufferedChannel()

	if len(result) != 3 {
		t.Fatalf("got %d elements, want 3", len(result))
	}

	expected := []int{1, 2, 3}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("index %d: got %d, want %d", i, result[i], v)
		}
	}
}

func TestSumWithChannel(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{10, 20, 30}, 60},
		{[]int{}, 0},
	}

	for _, tc := range tests {
		result := SumWithChannel(tc.input)
		if result != tc.expected {
			t.Errorf("SumWithChannel(%v): got %d, want %d", tc.input, result, tc.expected)
		}
	}
}

func TestCollectFromChannel(t *testing.T) {
	result := CollectFromChannel(5)
	expected := []int{0, 1, 2, 3, 4}

	if len(result) != len(expected) {
		t.Fatalf("got %d elements, want %d", len(result), len(expected))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("index %d: got %d, want %d", i, result[i], v)
		}
	}
}

func TestSelectFirst(t *testing.T) {
	// Test ch1 first
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "first"

	result := SelectFirst(ch1, ch2)
	if result != "first" {
		t.Errorf("got %q, want %q", result, "first")
	}

	// Test ch2 first
	ch1 = make(chan string, 1)
	ch2 = make(chan string, 1)
	ch2 <- "second"

	result = SelectFirst(ch1, ch2)
	if result != "second" {
		t.Errorf("got %q, want %q", result, "second")
	}
}

func TestWithTimeoutSuccess(t *testing.T) {
	fast := func() int {
		time.Sleep(10 * time.Millisecond)
		return 42
	}

	result, ok := WithTimeout(fast, 100*time.Millisecond)
	if !ok {
		t.Error("expected success, got timeout")
	}
	if result != 42 {
		t.Errorf("got %d, want 42", result)
	}
}

func TestWithTimeoutFailure(t *testing.T) {
	slow := func() int {
		time.Sleep(200 * time.Millisecond)
		return 42
	}

	_, ok := WithTimeout(slow, 50*time.Millisecond)
	if ok {
		t.Error("expected timeout, got success")
	}
}

func TestSumParallel(t *testing.T) {
	slices := [][]int{
		{1, 2, 3},       // 6
		{4, 5, 6},       // 15
		{7, 8, 9},       // 24
		{10, 11, 12},    // 33
	}

	result := SumParallel(slices)
	expected := 78 // 6 + 15 + 24 + 33

	if result != expected {
		t.Errorf("got %d, want %d", result, expected)
	}
}

func TestWorkerPool(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}
	result := WorkerPool(jobs, 3)

	if len(result) != len(jobs) {
		t.Fatalf("got %d results, want %d", len(result), len(jobs))
	}

	// Sort both for comparison (order doesn't matter)
	expected := []int{1, 4, 9, 16, 25}
	sort.Ints(result)
	sort.Ints(expected)

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("got %v, want %v", result, expected)
			break
		}
	}
}

func TestFanOutFanIn(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	// Each doubled: 2, 4, 6, 8, 10 = 30
	result := FanOutFanIn(nums, 3)

	if result != 30 {
		t.Errorf("got %d, want 30", result)
	}
}

func TestCounter(t *testing.T) {
	c := &Counter{}

	c.Increment()
	c.Increment()
	c.Increment()

	if c.Value() != 3 {
		t.Errorf("got %d, want 3", c.Value())
	}
}

func TestConcurrentIncrement(t *testing.T) {
	c := &Counter{}
	ConcurrentIncrement(c, 100)

	if c.Value() != 100 {
		t.Errorf("got %d, want 100 (race condition?)", c.Value())
	}
}

func TestConcurrentIncrementRaceDetection(t *testing.T) {
	// This test verifies mutex is working
	// Run with: go test -race
	c := &Counter{}
	ConcurrentIncrement(c, 1000)

	if c.Value() != 1000 {
		t.Errorf("got %d, want 1000", c.Value())
	}
}
