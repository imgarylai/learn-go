package collections

// Exercise 4: Slices and Maps
//
// Go's slices are like JS arrays, maps are like JS objects/Map.
// No built-in map/filter/reduce - you write loops!
// Run tests with: go test -v

// 1. Create and populate a slice
// In JS: const nums = [1, 2, 3]; nums.push(4, 5);
func CreateSlice() []int {
	// TODO: create slice with [1, 2, 3], append 4 and 5, return it
	return nil
}

// 2. Get a sub-slice (like JS array.slice())
// In JS: nums.slice(1, 3)
func SliceMiddle(nums []int) []int {
	// TODO: return elements from index 1 to 3 (exclusive)
	// If slice has less than 3 elements, return empty slice
	return nil
}

// 3. Double each element (like JS map)
// In JS: nums.map(n => n * 2)
func Double(nums []int) []int {
	// TODO: return new slice with each element doubled
	return nil
}

// 4. Filter elements (like JS filter)
// In JS: nums.filter(n => n > threshold)
func FilterGreaterThan(nums []int, threshold int) []int {
	// TODO: return only numbers greater than threshold
	return nil
}

// 5. Sum all elements (like JS reduce)
// In JS: nums.reduce((sum, n) => sum + n, 0)
func Sum(nums []int) int {
	// TODO: return sum of all numbers
	return 0
}

// 6. Find maximum value
// In JS: Math.max(...nums)
func Max(nums []int) int {
	// TODO: return the maximum value
	// If slice is empty, return 0
	return 0
}

// 7. Create a map (like JS object or Map)
// In JS: const scores = { alice: 95, bob: 87, charlie: 92 };
func CreateScores() map[string]int {
	// TODO: create and return map with alice:95, bob:87, charlie:92
	return nil
}

// 8. Get value from map with existence check
// In JS: scores.hasOwnProperty("alice") ? scores.alice : defaultVal
func GetScore(scores map[string]int, name string) (int, bool) {
	// TODO: return score and whether name exists
	// Hint: value, ok := map[key]
	return 0, false
}

// 9. Find the key with highest value
// In JS: Object.entries(scores).reduce((a, b) => a[1] > b[1] ? a : b)[0]
func GetTopScorer(scores map[string]int) string {
	// TODO: return name of person with highest score
	// If map is empty, return ""
	return ""
}

// 10. Delete from map
// In JS: delete scores.bob
func RemovePlayer(scores map[string]int, name string) {
	// TODO: remove the player from the map
}

// 11. Count occurrences
// In JS: arr.reduce((acc, x) => { acc[x] = (acc[x] || 0) + 1; return acc; }, {})
func CountOccurrences(items []string) map[string]int {
	// TODO: count how many times each item appears
	return nil
}

// Person for struct slice exercises
type Person struct {
	Name string
	Age  int
}

// 12. Filter slice of structs
// In JS: people.filter(p => p.age >= 18)
func GetAdults(people []Person) []Person {
	// TODO: return only people with Age >= 18
	return nil
}

// 13. Extract field from structs (like JS map)
// In JS: people.map(p => p.name)
func GetNames(people []Person) []string {
	// TODO: return slice of all names
	return nil
}

// 14. Find by field value
// In JS: people.find(p => p.name === name)
func FindByName(people []Person, name string) *Person {
	// TODO: return pointer to person with matching name, or nil if not found
	return nil
}
