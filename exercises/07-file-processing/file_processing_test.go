package fileprocessing

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// Helper to create temp directory for tests
func setupTestDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "fileprocessing-test-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	return dir
}

// Helper to write test file
func writeTestFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestReadLines(t *testing.T) {
	dir := setupTestDir(t)
	path := writeTestFile(t, dir, "test.txt", "line1\nline2\nline3")

	lines, err := ReadLines(path)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}
	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("got %v, want %v", lines, expected)
	}
}

func TestReadLinesEmpty(t *testing.T) {
	dir := setupTestDir(t)
	path := writeTestFile(t, dir, "empty.txt", "")

	lines, err := ReadLines(path)
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	if len(lines) != 0 {
		t.Errorf("expected empty slice, got %v", lines)
	}
}

func TestReadLinesNotFound(t *testing.T) {
	_, err := ReadLines("nonexistent.txt")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}

func TestWriteLines(t *testing.T) {
	dir := setupTestDir(t)
	path := filepath.Join(dir, "output.txt")

	lines := []string{"hello", "world", "go"}
	if err := WriteLines(path, lines); err != nil {
		t.Fatalf("WriteLines failed: %v", err)
	}

	// Verify
	readBack, _ := ReadLines(path)
	if !reflect.DeepEqual(readBack, lines) {
		t.Errorf("got %v, want %v", readBack, lines)
	}
}

func TestCountLines(t *testing.T) {
	dir := setupTestDir(t)
	path := writeTestFile(t, dir, "count.txt", "one\ntwo\nthree\nfour\nfive")

	count, err := CountLines(path)
	if err != nil {
		t.Fatalf("CountLines failed: %v", err)
	}

	if count != 5 {
		t.Errorf("got %d, want 5", count)
	}
}

func TestReadCSV(t *testing.T) {
	dir := setupTestDir(t)
	csvContent := `name,age,email
Alice,30,alice@example.com
Bob,25,bob@example.com`
	path := writeTestFile(t, dir, "people.csv", csvContent)

	people, err := ReadCSV(path)
	if err != nil {
		t.Fatalf("ReadCSV failed: %v", err)
	}

	expected := []Person{
		{Name: "Alice", Age: 30, Email: "alice@example.com"},
		{Name: "Bob", Age: 25, Email: "bob@example.com"},
	}

	if !reflect.DeepEqual(people, expected) {
		t.Errorf("got %+v, want %+v", people, expected)
	}
}

func TestWriteCSV(t *testing.T) {
	dir := setupTestDir(t)
	path := filepath.Join(dir, "output.csv")

	people := []Person{
		{Name: "Charlie", Age: 35, Email: "charlie@test.com"},
		{Name: "Diana", Age: 28, Email: "diana@test.com"},
	}

	if err := WriteCSV(path, people); err != nil {
		t.Fatalf("WriteCSV failed: %v", err)
	}

	// Read back and verify
	readBack, err := ReadCSV(path)
	if err != nil {
		t.Fatalf("ReadCSV failed: %v", err)
	}

	if !reflect.DeepEqual(readBack, people) {
		t.Errorf("got %+v, want %+v", readBack, people)
	}
}

func TestFilterCSV(t *testing.T) {
	dir := setupTestDir(t)
	inputCSV := `name,age,email
Alice,30,alice@example.com
Bob,17,bob@example.com
Charlie,45,charlie@example.com
Diana,15,diana@example.com`
	inputPath := writeTestFile(t, dir, "input.csv", inputCSV)
	outputPath := filepath.Join(dir, "filtered.csv")

	if err := FilterCSV(inputPath, outputPath, 18); err != nil {
		t.Fatalf("FilterCSV failed: %v", err)
	}

	filtered, err := ReadCSV(outputPath)
	if err != nil {
		t.Fatalf("ReadCSV failed: %v", err)
	}

	if len(filtered) != 2 {
		t.Errorf("expected 2 adults, got %d", len(filtered))
	}

	for _, p := range filtered {
		if p.Age < 18 {
			t.Errorf("found person under 18: %+v", p)
		}
	}
}

func TestReadJSON(t *testing.T) {
	dir := setupTestDir(t)
	jsonContent := `[
		{"name": "Alice", "age": 30, "email": "alice@example.com"},
		{"name": "Bob", "age": 25, "email": "bob@example.com"}
	]`
	path := writeTestFile(t, dir, "people.json", jsonContent)

	people, err := ReadJSON(path)
	if err != nil {
		t.Fatalf("ReadJSON failed: %v", err)
	}

	if len(people) != 2 {
		t.Errorf("expected 2 people, got %d", len(people))
	}

	if people[0].Name != "Alice" || people[0].Age != 30 {
		t.Errorf("unexpected first person: %+v", people[0])
	}
}

func TestWriteJSON(t *testing.T) {
	dir := setupTestDir(t)
	path := filepath.Join(dir, "output.json")

	people := []Person{
		{Name: "Eve", Age: 22, Email: "eve@test.com"},
	}

	if err := WriteJSON(path, people); err != nil {
		t.Fatalf("WriteJSON failed: %v", err)
	}

	// Read back
	readBack, err := ReadJSON(path)
	if err != nil {
		t.Fatalf("ReadJSON failed: %v", err)
	}

	if !reflect.DeepEqual(readBack, people) {
		t.Errorf("got %+v, want %+v", readBack, people)
	}
}

func TestConvertCSVToJSON(t *testing.T) {
	dir := setupTestDir(t)
	csvContent := `name,age,email
Frank,40,frank@example.com`
	csvPath := writeTestFile(t, dir, "convert.csv", csvContent)
	jsonPath := filepath.Join(dir, "convert.json")

	if err := ConvertCSVToJSON(csvPath, jsonPath); err != nil {
		t.Fatalf("ConvertCSVToJSON failed: %v", err)
	}

	people, err := ReadJSON(jsonPath)
	if err != nil {
		t.Fatalf("ReadJSON failed: %v", err)
	}

	if len(people) != 1 || people[0].Name != "Frank" {
		t.Errorf("unexpected result: %+v", people)
	}
}

func TestProcessLargeFile(t *testing.T) {
	dir := setupTestDir(t)
	content := "line1\nline2\nline3\nline4\nline5"
	path := writeTestFile(t, dir, "large.txt", content)

	var lines []string
	var lineNums []int

	err := ProcessLargeFile(path, func(lineNum int, line string) error {
		lineNums = append(lineNums, lineNum)
		lines = append(lines, line)
		return nil
	})

	if err != nil {
		t.Fatalf("ProcessLargeFile failed: %v", err)
	}

	if len(lines) != 5 {
		t.Errorf("expected 5 lines, got %d", len(lines))
	}

	expectedNums := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(lineNums, expectedNums) {
		t.Errorf("line numbers: got %v, want %v", lineNums, expectedNums)
	}
}

// ============ Tests using real CSV files from testdata/ ============

func TestReadLinesFromTestdata(t *testing.T) {
	lines, err := ReadLines("testdata/sample.txt")
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	if len(lines) != 5 {
		t.Errorf("expected 5 lines, got %d", len(lines))
	}

	if lines[0] != "Hello, Go!" {
		t.Errorf("first line: got %q, want %q", lines[0], "Hello, Go!")
	}
}

func TestReadCSVFromTestdata(t *testing.T) {
	people, err := ReadCSV("testdata/people.csv")
	if err != nil {
		t.Fatalf("ReadCSV failed: %v", err)
	}

	if len(people) != 5 {
		t.Errorf("expected 5 people, got %d", len(people))
	}

	// Check first person
	if people[0].Name != "Alice" || people[0].Age != 30 {
		t.Errorf("first person: got %+v, want Alice/30", people[0])
	}
}

func TestReadProducts(t *testing.T) {
	products, err := ReadProducts("testdata/products.csv")
	if err != nil {
		t.Fatalf("ReadProducts failed: %v", err)
	}

	if len(products) != 8 {
		t.Errorf("expected 8 products, got %d", len(products))
	}

	// Check first product
	if products[0].ID != 1 || products[0].Name != "Laptop" {
		t.Errorf("first product: got %+v", products[0])
	}

	if products[0].Price != 999.99 {
		t.Errorf("laptop price: got %f, want 999.99", products[0].Price)
	}
}

func TestFilterProductsByCategory(t *testing.T) {
	products, _ := ReadProducts("testdata/products.csv")
	electronics := FilterProductsByCategory(products, "Electronics")

	if len(electronics) != 3 {
		t.Errorf("expected 3 electronics, got %d", len(electronics))
	}

	for _, p := range electronics {
		if p.Category != "Electronics" {
			t.Errorf("found non-electronics: %+v", p)
		}
	}
}

func TestCalculateTotalValue(t *testing.T) {
	products, _ := ReadProducts("testdata/products.csv")
	total := CalculateTotalValue(products)

	// 999.99 + 79.99 + 12.99 + 4.99 + 49.99 + 29.99 + 19.99 + 34.99 = 1232.92
	expected := 1232.92
	if total < expected-0.01 || total > expected+0.01 {
		t.Errorf("total: got %f, want %f", total, expected)
	}
}

func TestFindMostExpensive(t *testing.T) {
	products, _ := ReadProducts("testdata/products.csv")
	most := FindMostExpensive(products)

	if most == nil {
		t.Fatal("FindMostExpensive returned nil")
	}

	if most.Name != "Laptop" {
		t.Errorf("most expensive: got %s, want Laptop", most.Name)
	}
}

func TestFindMostExpensiveEmpty(t *testing.T) {
	most := FindMostExpensive([]Product{})
	if most != nil {
		t.Errorf("expected nil for empty slice, got %+v", most)
	}
}

func TestGroupProductsByCategory(t *testing.T) {
	products, _ := ReadProducts("testdata/products.csv")
	grouped := GroupProductsByCategory(products)

	if len(grouped) != 4 {
		t.Errorf("expected 4 categories, got %d", len(grouped))
	}

	if len(grouped["Electronics"]) != 3 {
		t.Errorf("Electronics: expected 3, got %d", len(grouped["Electronics"]))
	}

	if len(grouped["Kitchen"]) != 2 {
		t.Errorf("Kitchen: expected 2, got %d", len(grouped["Kitchen"]))
	}
}
