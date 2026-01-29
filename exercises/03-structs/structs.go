package structs

// Exercise 3: Structs and Methods
//
// Go doesn't have classes, but structs + methods give you similar power.
// Think of it as: class = struct + methods
// Run tests with: go test -v

import "fmt"

// User represents a user (like a TS interface or class)
// In TS: interface User { id: number; name: string; email: string; }
type User struct {
	ID    int
	Name  string
	Email string
}

// 1. Constructor function - Go convention: NewXxx
// In JS: constructor(id, name, email) { this.id = id; ... }
func NewUser(id int, name, email string) *User {
	// TODO: return a pointer to a new User
	return nil
}

// 2. Method with value receiver - doesn't modify original
// In JS: getDisplayName() { return `${this.name} <${this.email}>`; }
func (u User) DisplayName() string {
	// TODO: return "Name <email>" format
	return ""
}

// 3. Method with pointer receiver - CAN modify the struct
// In JS: updateEmail(newEmail) { this.email = newEmail; }
func (u *User) UpdateEmail(newEmail string) {
	// TODO: update the user's email
}

// 4. Method that checks something
func (u User) IsValidEmail() bool {
	// TODO: return true if email contains "@"
	// Hint: use strings.Contains or just loop through
	return false
}

// Admin embeds User (like inheritance/composition)
// In JS: class Admin extends User { role: string; }
type Admin struct {
	User // embedded - Admin "inherits" User's fields and methods
	Role string
}

// 5. Constructor for embedded struct
func NewAdmin(id int, name, email, role string) *Admin {
	// TODO: return a new Admin with the given values
	return nil
}

// 6. Method on embedded struct (Admin gets User methods for free!)
// This is an ADDITIONAL method specific to Admin
func (a Admin) CanDelete() bool {
	// TODO: return true if role is "superadmin"
	return false
}

// Product with struct tags for JSON serialization
// In TS: decorators or runtime metadata
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 7. Constructor for Product
func NewProduct(id int, name string, price float64) Product {
	// TODO: return a new Product (not pointer - value type)
	return Product{}
}

// 8. Method to apply discount
func (p Product) WithDiscount(percent float64) Product {
	// TODO: return NEW product with discounted price
	// Don't modify original - return a copy
	// Example: 20% discount on $100 = $80
	return Product{}
}

// Rectangle for area/perimeter calculations
type Rectangle struct {
	Width  float64
	Height float64
}

// 9. Calculate area
func (r Rectangle) Area() float64 {
	// TODO: return width * height
	return 0
}

// 10. Calculate perimeter
func (r Rectangle) Perimeter() float64 {
	// TODO: return 2 * (width + height)
	return 0
}

// Keep import used
var _ = fmt.Sprintf
