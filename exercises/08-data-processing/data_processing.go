package dataprocessing

// Exercise 8: Data Processing
//
// Practice data manipulation with slices, generics, and gota DataFrame.
// Run tests with: go test -v
//
// First, install gota:
//   go get github.com/go-gota/gota/dataframe
//   go get github.com/go-gota/gota/series

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// ============ Part 1: Pure Go (no external deps) ============

// Sale represents a sales record
type Sale struct {
	Product  string
	Quantity int
	Price    float64
	Region   string
}

// 1. Filter - return sales where quantity > minQty
// In Python: df[df['quantity'] > min_qty]
func FilterSales(sales []Sale, minQty int) []Sale {
	// TODO: filter and return matching sales
	return nil
}

// 2. Map - extract all product names
// In Python: df['product'].tolist()
func GetProductNames(sales []Sale) []string {
	// TODO: return slice of product names
	return nil
}

// 3. Reduce - calculate total revenue (quantity * price for all sales)
// In Python: (df['quantity'] * df['price']).sum()
func TotalRevenue(sales []Sale) float64 {
	// TODO: sum of quantity * price for each sale
	return 0
}

// 4. GroupBy - group sales by region, return map of region -> []Sale
// In Python: df.groupby('region')
func GroupByRegion(sales []Sale) map[string][]Sale {
	// TODO: group sales by region
	return nil
}

// 5. Aggregate - calculate total revenue per region
// In Python: df.groupby('region').apply(lambda x: (x['quantity'] * x['price']).sum())
func RevenueByRegion(sales []Sale) map[string]float64 {
	// TODO: total revenue for each region
	return nil
}

// 6. TopN - return top N sales by revenue (quantity * price)
// In Python: df.nlargest(n, 'revenue')
func TopNSales(sales []Sale, n int) []Sale {
	// TODO: sort by revenue descending, return top N
	// Hint: use sort.Slice
	return nil
}

// 7. Unique - return unique product names
// In Python: df['product'].unique()
func UniqueProducts(sales []Sale) []string {
	// TODO: return unique product names
	// Hint: use a map to track seen values
	return nil
}

// 8. CountBy - count sales per product
// In Python: df['product'].value_counts()
func SalesCountByProduct(sales []Sale) map[string]int {
	// TODO: count occurrences of each product
	return nil
}

// ============ Part 2: Generic helpers (reusable) ============

// 9. Generic Filter - works with any type
// In Python: list(filter(predicate, items))
func Filter[T any](items []T, predicate func(T) bool) []T {
	// TODO: return items where predicate returns true
	return nil
}

// 10. Generic Map - transform items
// In Python: list(map(transform, items))
func Map[T, U any](items []T, transform func(T) U) []U {
	// TODO: apply transform to each item
	return nil
}

// 11. Generic Reduce - fold items into single value
// In Python: functools.reduce(reducer, items, initial)
func Reduce[T, U any](items []T, initial U, reducer func(U, T) U) U {
	// TODO: reduce items to single value
	return initial
}

// 12. Generic GroupBy
func GroupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
	// TODO: group items by key function
	return nil
}

// ============ Part 3: Gota DataFrame ============

// 13. Create DataFrame from sales slice
// In Python: pd.DataFrame(sales)
func SalesToDataFrame(sales []Sale) dataframe.DataFrame {
	// TODO: use dataframe.LoadStructs
	return dataframe.DataFrame{}
}

// 14. Filter DataFrame - sales with quantity > minQty
// In Python: df[df['Quantity'] > min_qty]
func FilterDataFrame(df dataframe.DataFrame, minQty int) dataframe.DataFrame {
	// TODO: use df.Filter with dataframe.F
	return dataframe.DataFrame{}
}

// 15. Select columns from DataFrame
// In Python: df[['Product', 'Price']]
func SelectColumns(df dataframe.DataFrame, cols ...string) dataframe.DataFrame {
	// TODO: use df.Select
	return dataframe.DataFrame{}
}

// 16. Sort DataFrame by column
// In Python: df.sort_values('Quantity', ascending=False)
func SortByQuantity(df dataframe.DataFrame, descending bool) dataframe.DataFrame {
	// TODO: use df.Arrange with dataframe.Sort or dataframe.RevSort
	return dataframe.DataFrame{}
}

// 17. Get column statistics
// In Python: df['Quantity'].mean(), df['Quantity'].sum()
type ColumnStats struct {
	Sum  float64
	Mean float64
	Min  float64
	Max  float64
}

func GetQuantityStats(df dataframe.DataFrame) ColumnStats {
	// TODO: get statistics from Quantity column
	// Hint: df.Col("Quantity") returns a series.Series
	return ColumnStats{}
}

// ============ Part 4: Working with Real CSV Files ============
// Use the CSV files in testdata/ folder

// Employee represents an employee from employees.csv
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     int
	Years      int
}

// 18. ReadEmployees reads employees.csv from testdata folder
func ReadEmployees(filename string) ([]Employee, error) {
	// TODO: Read CSV and parse into []Employee
	return nil, nil
}

// 19. AverageSalaryByDepartment calculates avg salary per department
// In Python: df.groupby('department')['salary'].mean()
func AverageSalaryByDepartment(employees []Employee) map[string]float64 {
	// TODO: Return map of department -> average salary
	return nil
}

// 20. TopEarners returns top N employees by salary
func TopEarners(employees []Employee, n int) []Employee {
	// TODO: Sort by salary descending, return top N
	return nil
}

// 21. FilterByExperience returns employees with >= minYears
func FilterByExperience(employees []Employee, minYears int) []Employee {
	// TODO: Filter employees by years of experience
	return nil
}

// 22. TotalPayroll calculates sum of all salaries
func TotalPayroll(employees []Employee) int {
	// TODO: Sum all salaries
	return 0
}

// 23. ReadSalesCSV reads sales.csv and returns []Sale
func ReadSalesCSV(filename string) ([]Sale, error) {
	// TODO: Read sales.csv and parse into []Sale
	return nil, nil
}

// Keep imports used
var (
	_ = sort.Slice
	_ = dataframe.DataFrame{}
	_ = series.Series{}
	_ = csv.Reader{}
	_ = os.Open
	_ = strconv.Atoi
)
