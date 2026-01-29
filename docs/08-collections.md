# Slices and Maps in Go

Coming from JavaScript arrays and objects? Here's how Go collections work.

## Arrays vs Slices

Go has both arrays (fixed size) and slices (dynamic). You'll almost always use **slices**.

```go
// Array - fixed size, rarely used directly
var arr [5]int                    // [0, 0, 0, 0, 0]
arr := [3]string{"a", "b", "c"}   // Size is part of type

// Slice - dynamic, what you'll use 99% of the time
var slice []int                   // nil slice
slice := []string{"a", "b", "c"}  // Slice literal
slice := make([]int, 5)           // [0, 0, 0, 0, 0] with len=5
slice := make([]int, 0, 10)       // Empty with capacity 10
```

### JS Comparison

```javascript
// JavaScript - arrays are always dynamic
const arr = [1, 2, 3];
arr.push(4);  // [1, 2, 3, 4]
```

```go
// Go - slices are dynamic
slice := []int{1, 2, 3}
slice = append(slice, 4)  // [1, 2, 3, 4] - must reassign!
```

## Slice Operations

### Creating Slices

```go
// Literal
nums := []int{1, 2, 3, 4, 5}

// make(type, length, capacity)
slice := make([]int, 3)      // [0, 0, 0] len=3, cap=3
slice := make([]int, 3, 10)  // [0, 0, 0] len=3, cap=10

// From array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2, 3, 4] - shares memory with arr!
```

### Length and Capacity

```go
slice := make([]int, 3, 10)
fmt.Println(len(slice))  // 3 - current length
fmt.Println(cap(slice))  // 10 - capacity before reallocation
```

### Append

```go
slice := []int{1, 2, 3}

// MUST reassign - append may return new underlying array
slice = append(slice, 4)           // [1, 2, 3, 4]
slice = append(slice, 5, 6, 7)     // [1, 2, 3, 4, 5, 6, 7]

// Append another slice (spread operator equivalent)
other := []int{8, 9}
slice = append(slice, other...)    // Note the ...
```

### Slicing (like JS slice)

```go
nums := []int{0, 1, 2, 3, 4, 5}

nums[1:4]   // [1, 2, 3]     - from index 1 to 3
nums[:3]    // [0, 1, 2]     - first 3 elements
nums[3:]    // [3, 4, 5]     - from index 3 to end
nums[:]     // [0, 1, 2, 3, 4, 5] - copy of entire slice
```

### Copy

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)  // dst is now [1, 2, 3], independent copy
```

### Delete Element

```go
// No built-in delete for slices - use append trick
slice := []int{1, 2, 3, 4, 5}
i := 2  // index to remove

// Remove element at index i (preserves order)
slice = append(slice[:i], slice[i+1:]...)  // [1, 2, 4, 5]
```

## Iterating Slices

```go
names := []string{"Alice", "Bob", "Charlie"}

// Index and value
for i, name := range names {
    fmt.Printf("%d: %s\n", i, name)
}

// Value only
for _, name := range names {
    fmt.Println(name)
}

// Index only
for i := range names {
    fmt.Println(i)
}

// Classic for loop
for i := 0; i < len(names); i++ {
    fmt.Println(names[i])
}
```

### JS Comparison

```javascript
// JavaScript
names.forEach((name, i) => console.log(i, name));
for (const name of names) { }
for (const [i, name] of names.entries()) { }
```

## Maps

Maps are Go's equivalent to JavaScript objects/Map.

### Creating Maps

```go
// Literal
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}

// make
ages := make(map[string]int)
ages["Alice"] = 30
ages["Bob"] = 25

// nil map - can read but NOT write!
var ages map[string]int  // nil
_ = ages["Alice"]        // OK, returns zero value (0)
ages["Bob"] = 25         // PANIC! assignment to nil map
```

### JS Comparison

```javascript
// JavaScript object
const ages = {
    Alice: 30,
    Bob: 25,
};
ages.Charlie = 35;
ages["Diana"] = 28;

// JavaScript Map
const map = new Map([["Alice", 30], ["Bob", 25]]);
map.set("Charlie", 35);
```

```go
// Go map
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}
ages["Charlie"] = 35
```

## Map Operations

### Get and Set

```go
ages := map[string]int{"Alice": 30}

// Get
age := ages["Alice"]     // 30
age := ages["Unknown"]   // 0 (zero value, NOT error)

// Check if key exists
age, ok := ages["Alice"]
if ok {
    fmt.Println("Found:", age)
}

// Common pattern - get with existence check
if age, ok := ages["Bob"]; ok {
    fmt.Println(age)
} else {
    fmt.Println("Bob not found")
}

// Set
ages["Bob"] = 25
```

### Delete

```go
delete(ages, "Alice")  // Remove key
delete(ages, "Unknown") // No error if key doesn't exist
```

### Length

```go
len(ages)  // Number of key-value pairs
```

## Iterating Maps

```go
ages := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}

// Key and value
for name, age := range ages {
    fmt.Printf("%s is %d\n", name, age)
}

// Keys only
for name := range ages {
    fmt.Println(name)
}

// Values only (rare)
for _, age := range ages {
    fmt.Println(age)
}
```

**Warning**: Map iteration order is **random**! Not insertion order like JS Map.

```go
// To iterate in order, sort keys first
keys := make([]string, 0, len(ages))
for k := range ages {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Printf("%s: %d\n", k, ages[k])
}
```

## Common Patterns

### Set (using map)

```go
// Go has no built-in Set, use map[T]bool or map[T]struct{}
seen := make(map[string]bool)
seen["apple"] = true
seen["banana"] = true

if seen["apple"] {
    fmt.Println("apple exists")
}

// Memory-efficient set using empty struct
seen := make(map[string]struct{})
seen["apple"] = struct{}{}
if _, ok := seen["apple"]; ok {
    fmt.Println("apple exists")
}
```

### Counting (like JS reduce)

```go
words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

counts := make(map[string]int)
for _, word := range words {
    counts[word]++  // Zero value is 0, so this works
}
// map[apple:3 banana:2 cherry:1]
```

### Grouping

```go
type Person struct {
    Name string
    City string
}

people := []Person{
    {"Alice", "NYC"},
    {"Bob", "LA"},
    {"Charlie", "NYC"},
}

byCity := make(map[string][]Person)
for _, p := range people {
    byCity[p.City] = append(byCity[p.City], p)
}
// map[NYC:[{Alice NYC} {Charlie NYC}] LA:[{Bob LA}]]
```

### Filter Slice

```go
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Filter even numbers
var evens []int
for _, n := range nums {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}
// [2, 4, 6, 8, 10]
```

### Map/Transform Slice

```go
nums := []int{1, 2, 3, 4, 5}

// Double each number
doubled := make([]int, len(nums))
for i, n := range nums {
    doubled[i] = n * 2
}
// [2, 4, 6, 8, 10]
```

## Slices Are Reference Types

Slices point to an underlying array. Be careful with modifications!

```go
original := []int{1, 2, 3, 4, 5}
slice := original[1:4]  // [2, 3, 4]

slice[0] = 999
fmt.Println(original)  // [1, 999, 3, 4, 5] - modified!
```

### Safe Copy

```go
original := []int{1, 2, 3, 4, 5}
safeCopy := make([]int, len(original))
copy(safeCopy, original)

safeCopy[0] = 999
fmt.Println(original)  // [1, 2, 3, 4, 5] - unchanged
```

## nil vs Empty

```go
// nil slice
var nilSlice []int
fmt.Println(nilSlice == nil)  // true
fmt.Println(len(nilSlice))    // 0

// Empty slice
emptySlice := []int{}
fmt.Println(emptySlice == nil) // false
fmt.Println(len(emptySlice))   // 0

// Both work the same with append
nilSlice = append(nilSlice, 1)     // Works!
emptySlice = append(emptySlice, 1) // Works!

// nil map - different behavior!
var nilMap map[string]int
_ = nilMap["key"]       // OK, returns 0
nilMap["key"] = 1       // PANIC!

// Empty map - safe
emptyMap := map[string]int{}
emptyMap["key"] = 1     // OK
```

## Performance Tips

### Pre-allocate with make

```go
// Bad - many reallocations
var result []int
for i := 0; i < 10000; i++ {
    result = append(result, i)
}

// Good - single allocation
result := make([]int, 0, 10000)
for i := 0; i < 10000; i++ {
    result = append(result, i)
}

// Best - if you know exact size
result := make([]int, 10000)
for i := 0; i < 10000; i++ {
    result[i] = i
}
```

### Map Size Hint

```go
// Pre-allocate map capacity
users := make(map[string]User, 1000)  // Hint: ~1000 entries
```

## Quick Reference

| Operation | Slice | Map |
|-----------|-------|-----|
| Create empty | `make([]T, 0)` or `[]T{}` | `make(map[K]V)` or `map[K]V{}` |
| Create with size | `make([]T, len, cap)` | `make(map[K]V, size)` |
| Length | `len(s)` | `len(m)` |
| Add | `s = append(s, v)` | `m[k] = v` |
| Get | `s[i]` | `m[k]` or `v, ok := m[k]` |
| Delete | `s = append(s[:i], s[i+1:]...)` | `delete(m, k)` |
| Contains | Loop to check | `_, ok := m[k]` |
| Iterate | `for i, v := range s` | `for k, v := range m` |
| Copy | `copy(dst, src)` | Loop or maps.Clone (Go 1.21+) |

## JS to Go Cheatsheet

| JavaScript | Go |
|------------|-----|
| `arr.push(x)` | `slice = append(slice, x)` |
| `arr.pop()` | `slice = slice[:len(slice)-1]` |
| `arr.shift()` | `slice = slice[1:]` |
| `arr.slice(1, 3)` | `slice[1:3]` |
| `arr.length` | `len(slice)` |
| `arr.map(fn)` | for loop |
| `arr.filter(fn)` | for loop |
| `arr.find(fn)` | for loop |
| `arr.includes(x)` | for loop or use map |
| `[...arr1, ...arr2]` | `append(arr1, arr2...)` |
| `obj.key` or `obj["key"]` | `m["key"]` |
| `delete obj.key` | `delete(m, "key")` |
| `Object.keys(obj)` | for loop to collect keys |
| `"key" in obj` | `_, ok := m["key"]` |
