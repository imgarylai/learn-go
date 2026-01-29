# Data Processing in Go

Coming from Python pandas? Go takes a different approach - no single dominant library, but powerful primitives.

## Go vs Pandas Philosophy

| Pandas | Go |
|--------|-----|
| One library does everything | Compose smaller tools |
| DataFrame as core abstraction | Slices of structs |
| Dynamic typing | Static typing |
| Interpreted (slower) | Compiled (faster) |
| Great for exploration | Great for production |

## Standard Library Approach

Go's philosophy: use slices of structs and write explicit loops.

### The "Go Way" - Slices of Structs

```python
# Python pandas
df = pd.DataFrame({'name': ['Alice', 'Bob'], 'age': [30, 25]})
adults = df[df['age'] >= 18]
avg_age = df['age'].mean()
```

```go
// Go - explicit but clear
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
}

// Filter
var adults []Person
for _, p := range people {
    if p.Age >= 18 {
        adults = append(adults, p)
    }
}

// Average
sum := 0
for _, p := range people {
    sum += p.Age
}
avg := float64(sum) / float64(len(people))
```

### Generic Helper Functions (Go 1.18+)

```go
// Filter any slice
func Filter[T any](slice []T, predicate func(T) bool) []T {
    result := make([]T, 0)
    for _, item := range slice {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

// Map any slice
func Map[T, U any](slice []T, transform func(T) U) []U {
    result := make([]U, len(slice))
    for i, item := range slice {
        result[i] = transform(item)
    }
    return result
}

// Reduce any slice
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
    result := initial
    for _, item := range slice {
        result = reducer(result, item)
    }
    return result
}

// Usage
adults := Filter(people, func(p Person) bool { return p.Age >= 18 })
names := Map(people, func(p Person) string { return p.Name })
totalAge := Reduce(people, 0, func(sum int, p Person) int { return sum + p.Age })
```

## Gota - Go's Pandas-like Library

[Gota](https://github.com/go-gota/gota) provides DataFrame and Series types.

```bash
go get github.com/go-gota/gota/dataframe
go get github.com/go-gota/gota/series
```

### Creating DataFrames

```go
import (
    "github.com/go-gota/gota/dataframe"
    "github.com/go-gota/gota/series"
)

// From structs
type Person struct {
    Name string
    Age  int
    City string
}

people := []Person{
    {"Alice", 30, "NYC"},
    {"Bob", 25, "LA"},
    {"Charlie", 35, "NYC"},
}

df := dataframe.LoadStructs(people)
fmt.Println(df)
```

### Reading CSV with Gota

```go
// From CSV file
file, _ := os.Open("data.csv")
df := dataframe.ReadCSV(file)

// From CSV string
csvStr := `Name,Age,City
Alice,30,NYC
Bob,25,LA`
df := dataframe.ReadCSV(strings.NewReader(csvStr))
```

### Basic Operations

```go
// Select columns
names := df.Select("Name")

// Filter rows
adults := df.Filter(
    dataframe.F{Colname: "Age", Comparator: series.GreaterEq, Comparando: 18},
)

// Multiple filters (AND)
nycAdults := df.Filter(
    dataframe.F{Colname: "Age", Comparator: series.GreaterEq, Comparando: 18},
    dataframe.F{Colname: "City", Comparator: series.Eq, Comparando: "NYC"},
)

// Sort
sorted := df.Arrange(dataframe.Sort("Age"))        // Ascending
sorted := df.Arrange(dataframe.RevSort("Age"))     // Descending

// First/Last N rows
first5 := df.Subset([]int{0, 1, 2, 3, 4})
```

### Aggregations

```go
// Describe (like pandas describe)
fmt.Println(df.Describe())

// Group by and aggregate
grouped := df.GroupBy("City").Aggregation(
    []dataframe.AggregationType{dataframe.Aggregation_MEAN},
    []string{"Age"},
)

// Column statistics
ages := df.Col("Age")
fmt.Println("Mean:", ages.Mean())
fmt.Println("Max:", ages.Max())
fmt.Println("Min:", ages.Min())
fmt.Println("StdDev:", ages.StdDev())
```

### Mutations

```go
// Add new column
df = df.Mutate(
    series.New([]string{"Junior", "Junior", "Senior"}, series.String, "Level"),
)

// Rename columns
df = df.Rename("FullName", "Name")

// Drop columns
df = df.Drop("City")
```

## When to Use What

| Use Case | Approach |
|----------|----------|
| Simple transformations | Slices + loops |
| Type-safe data pipelines | Slices of structs + generics |
| Exploratory analysis | Gota DataFrame |
| Production ETL | Slices of structs |
| Statistical analysis | Gonum |
| Large datasets | Slices + parallel processing |

## Numerical Computing with Gonum

[Gonum](https://www.gonum.org/) is Go's scientific computing library.

```bash
go get gonum.org/v1/gonum/...
```

```go
import (
    "gonum.org/v1/gonum/stat"
    "gonum.org/v1/gonum/floats"
)

data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

// Statistics
mean := stat.Mean(data, nil)
variance := stat.Variance(data, nil)
stdDev := stat.StdDev(data, nil)

// Operations
sum := floats.Sum(data)
floats.Scale(2.0, data)  // Multiply all by 2 (in-place)
```

## Parallel Data Processing

Go excels at parallel processing with goroutines.

```go
// Process data in parallel
func processInParallel[T, U any](items []T, workers int, process func(T) U) []U {
    results := make([]U, len(items))
    jobs := make(chan int, len(items))
    var wg sync.WaitGroup

    // Start workers
    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for i := range jobs {
                results[i] = process(items[i])
            }
        }()
    }

    // Send jobs
    for i := range items {
        jobs <- i
    }
    close(jobs)

    wg.Wait()
    return results
}

// Usage
squares := processInParallel([]int{1, 2, 3, 4, 5}, 4, func(n int) int {
    return n * n
})
```

## Common Data Processing Patterns

### Group By

```go
func GroupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
    result := make(map[K][]T)
    for _, item := range items {
        key := keyFn(item)
        result[key] = append(result[key], item)
    }
    return result
}

// Usage
byCity := GroupBy(people, func(p Person) string { return p.City })
```

### Unique Values

```go
func Unique[T comparable](items []T) []T {
    seen := make(map[T]bool)
    result := make([]T, 0)
    for _, item := range items {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }
    return result
}
```

### Count Occurrences

```go
func CountBy[T comparable](items []T) map[T]int {
    counts := make(map[T]int)
    for _, item := range items {
        counts[item]++
    }
    return counts
}
```

### Top N

```go
func TopN[T any](items []T, n int, less func(a, b T) bool) []T {
    sorted := make([]T, len(items))
    copy(sorted, items)
    sort.Slice(sorted, func(i, j int) bool {
        return less(sorted[j], sorted[i]) // Reverse for descending
    })
    if n > len(sorted) {
        n = len(sorted)
    }
    return sorted[:n]
}

// Usage
top3 := TopN(people, 3, func(a, b Person) bool { return a.Age < b.Age })
```

## Real-World Example: Sales Analysis

```go
type Sale struct {
    Date     time.Time
    Product  string
    Quantity int
    Price    float64
}

func analyzeSales(sales []Sale) {
    // Total revenue
    total := Reduce(sales, 0.0, func(sum float64, s Sale) float64 {
        return sum + float64(s.Quantity) * s.Price
    })

    // Revenue by product
    byProduct := GroupBy(sales, func(s Sale) string { return s.Product })
    productRevenue := make(map[string]float64)
    for product, productSales := range byProduct {
        productRevenue[product] = Reduce(productSales, 0.0, func(sum float64, s Sale) float64 {
            return sum + float64(s.Quantity) * s.Price
        })
    }

    // Top selling products
    type ProductSales struct {
        Product string
        Revenue float64
    }
    var rankings []ProductSales
    for p, r := range productRevenue {
        rankings = append(rankings, ProductSales{p, r})
    }
    sort.Slice(rankings, func(i, j int) bool {
        return rankings[i].Revenue > rankings[j].Revenue
    })
}
```

## Summary: Go vs Python for Data

| Aspect | Python/Pandas | Go |
|--------|---------------|-----|
| Speed | Slower (interpreted) | 10-100x faster |
| Memory | Higher overhead | Lower, predictable |
| Type safety | Runtime errors | Compile-time checks |
| Concurrency | GIL limitations | Native goroutines |
| Learning curve | Gentle | Steeper for data tasks |
| Best for | Exploration, notebooks | Production pipelines |

**Recommendation**: Use Python/pandas for exploration and prototyping, Go for production data pipelines where performance and reliability matter.
