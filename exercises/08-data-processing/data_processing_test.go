package dataprocessing

import (
	"reflect"
	"sort"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// Test data
func getSampleSales() []Sale {
	return []Sale{
		{Product: "Widget", Quantity: 10, Price: 25.0, Region: "North"},
		{Product: "Gadget", Quantity: 5, Price: 50.0, Region: "South"},
		{Product: "Widget", Quantity: 8, Price: 25.0, Region: "South"},
		{Product: "Gizmo", Quantity: 15, Price: 30.0, Region: "North"},
		{Product: "Gadget", Quantity: 3, Price: 50.0, Region: "East"},
	}
}

// ============ Part 1: Pure Go Tests ============

func TestFilterSales(t *testing.T) {
	sales := getSampleSales()
	filtered := FilterSales(sales, 7)

	if len(filtered) != 3 {
		t.Errorf("expected 3 sales with qty > 7, got %d", len(filtered))
	}

	for _, s := range filtered {
		if s.Quantity <= 7 {
			t.Errorf("found sale with qty %d, expected > 7", s.Quantity)
		}
	}
}

func TestGetProductNames(t *testing.T) {
	sales := getSampleSales()
	names := GetProductNames(sales)

	if len(names) != 5 {
		t.Errorf("expected 5 names, got %d", len(names))
	}

	expected := []string{"Widget", "Gadget", "Widget", "Gizmo", "Gadget"}
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("got %v, want %v", names, expected)
	}
}

func TestTotalRevenue(t *testing.T) {
	sales := getSampleSales()
	revenue := TotalRevenue(sales)

	// 10*25 + 5*50 + 8*25 + 15*30 + 3*50 = 250 + 250 + 200 + 450 + 150 = 1300
	expected := 1300.0
	if revenue != expected {
		t.Errorf("got %.2f, want %.2f", revenue, expected)
	}
}

func TestGroupByRegion(t *testing.T) {
	sales := getSampleSales()
	grouped := GroupByRegion(sales)

	if len(grouped) != 3 {
		t.Errorf("expected 3 regions, got %d", len(grouped))
	}

	if len(grouped["North"]) != 2 {
		t.Errorf("expected 2 North sales, got %d", len(grouped["North"]))
	}

	if len(grouped["South"]) != 2 {
		t.Errorf("expected 2 South sales, got %d", len(grouped["South"]))
	}

	if len(grouped["East"]) != 1 {
		t.Errorf("expected 1 East sale, got %d", len(grouped["East"]))
	}
}

func TestRevenueByRegion(t *testing.T) {
	sales := getSampleSales()
	revenue := RevenueByRegion(sales)

	// North: 10*25 + 15*30 = 700
	// South: 5*50 + 8*25 = 450
	// East: 3*50 = 150
	if revenue["North"] != 700 {
		t.Errorf("North: got %.2f, want 700", revenue["North"])
	}
	if revenue["South"] != 450 {
		t.Errorf("South: got %.2f, want 450", revenue["South"])
	}
	if revenue["East"] != 150 {
		t.Errorf("East: got %.2f, want 150", revenue["East"])
	}
}

func TestTopNSales(t *testing.T) {
	sales := getSampleSales()
	top2 := TopNSales(sales, 2)

	if len(top2) != 2 {
		t.Errorf("expected 2 sales, got %d", len(top2))
	}

	// Top 2 by revenue: Gizmo (450), Widget-North or Gadget-South (250 each)
	if top2[0].Product != "Gizmo" {
		t.Errorf("expected Gizmo as top, got %s", top2[0].Product)
	}
}

func TestUniqueProducts(t *testing.T) {
	sales := getSampleSales()
	unique := UniqueProducts(sales)

	if len(unique) != 3 {
		t.Errorf("expected 3 unique products, got %d", len(unique))
	}

	// Sort for comparison
	sort.Strings(unique)
	expected := []string{"Gadget", "Gizmo", "Widget"}
	if !reflect.DeepEqual(unique, expected) {
		t.Errorf("got %v, want %v", unique, expected)
	}
}

func TestSalesCountByProduct(t *testing.T) {
	sales := getSampleSales()
	counts := SalesCountByProduct(sales)

	if counts["Widget"] != 2 {
		t.Errorf("Widget: got %d, want 2", counts["Widget"])
	}
	if counts["Gadget"] != 2 {
		t.Errorf("Gadget: got %d, want 2", counts["Gadget"])
	}
	if counts["Gizmo"] != 1 {
		t.Errorf("Gizmo: got %d, want 1", counts["Gizmo"])
	}
}

// ============ Part 2: Generic Helpers Tests ============

func TestGenericFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })

	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(evens, expected) {
		t.Errorf("got %v, want %v", evens, expected)
	}
}

func TestGenericMap(t *testing.T) {
	nums := []int{1, 2, 3}
	doubled := Map(nums, func(n int) int { return n * 2 })

	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(doubled, expected) {
		t.Errorf("got %v, want %v", doubled, expected)
	}

	// Test type transformation
	strs := Map(nums, func(n int) string {
		return string(rune('A' + n - 1))
	})
	expectedStrs := []string{"A", "B", "C"}
	if !reflect.DeepEqual(strs, expectedStrs) {
		t.Errorf("got %v, want %v", strs, expectedStrs)
	}
}

func TestGenericReduce(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	sum := Reduce(nums, 0, func(acc int, n int) int { return acc + n })

	if sum != 15 {
		t.Errorf("got %d, want 15", sum)
	}

	// Test with different types
	words := []string{"hello", " ", "world"}
	concat := Reduce(words, "", func(acc string, s string) string { return acc + s })
	if concat != "hello world" {
		t.Errorf("got %q, want %q", concat, "hello world")
	}
}

func TestGenericGroupBy(t *testing.T) {
	type Item struct {
		Category string
		Value    int
	}
	items := []Item{
		{"A", 1}, {"B", 2}, {"A", 3}, {"B", 4}, {"C", 5},
	}

	grouped := GroupBy(items, func(i Item) string { return i.Category })

	if len(grouped) != 3 {
		t.Errorf("expected 3 groups, got %d", len(grouped))
	}
	if len(grouped["A"]) != 2 {
		t.Errorf("expected 2 items in A, got %d", len(grouped["A"]))
	}
}

// ============ Part 3: Gota DataFrame Tests ============

func TestSalesToDataFrame(t *testing.T) {
	sales := getSampleSales()
	df := SalesToDataFrame(sales)

	if df.Nrow() != 5 {
		t.Errorf("expected 5 rows, got %d", df.Nrow())
	}

	if df.Ncol() != 4 {
		t.Errorf("expected 4 columns, got %d", df.Ncol())
	}
}

func TestFilterDataFrame(t *testing.T) {
	sales := getSampleSales()
	df := SalesToDataFrame(sales)
	filtered := FilterDataFrame(df, 7)

	if filtered.Nrow() != 3 {
		t.Errorf("expected 3 rows after filter, got %d", filtered.Nrow())
	}
}

func TestSelectColumns(t *testing.T) {
	sales := getSampleSales()
	df := SalesToDataFrame(sales)
	selected := SelectColumns(df, "Product", "Price")

	if selected.Ncol() != 2 {
		t.Errorf("expected 2 columns, got %d", selected.Ncol())
	}

	names := selected.Names()
	if names[0] != "Product" || names[1] != "Price" {
		t.Errorf("unexpected columns: %v", names)
	}
}

func TestSortByQuantity(t *testing.T) {
	sales := getSampleSales()
	df := SalesToDataFrame(sales)

	// Sort descending
	sorted := SortByQuantity(df, true)

	// First row should have highest quantity (15)
	firstQty := sorted.Elem(0, 1).Int()
	if firstQty != 15 {
		t.Errorf("expected first quantity to be 15, got %d", firstQty)
	}

	// Sort ascending
	sortedAsc := SortByQuantity(df, false)
	firstQtyAsc := sortedAsc.Elem(0, 1).Int()
	if firstQtyAsc != 3 {
		t.Errorf("expected first quantity to be 3, got %d", firstQtyAsc)
	}
}

func TestGetQuantityStats(t *testing.T) {
	sales := getSampleSales()
	df := SalesToDataFrame(sales)
	stats := GetQuantityStats(df)

	// Quantities: 10, 5, 8, 15, 3
	// Sum: 41, Mean: 8.2, Min: 3, Max: 15
	if stats.Sum != 41 {
		t.Errorf("Sum: got %.2f, want 41", stats.Sum)
	}
	if stats.Mean != 8.2 {
		t.Errorf("Mean: got %.2f, want 8.2", stats.Mean)
	}
	if stats.Min != 3 {
		t.Errorf("Min: got %.2f, want 3", stats.Min)
	}
	if stats.Max != 15 {
		t.Errorf("Max: got %.2f, want 15", stats.Max)
	}
}

// Keep imports
var (
	_ = series.Int
	_ = dataframe.LoadStructs
)
