package structs

import (
	"math"
	"testing"
)

func TestNewUser(t *testing.T) {
	user := NewUser(1, "Alice", "alice@example.com")

	if user == nil {
		t.Fatal("NewUser returned nil")
	}

	if user.ID != 1 {
		t.Errorf("ID: got %d, want 1", user.ID)
	}
	if user.Name != "Alice" {
		t.Errorf("Name: got %q, want %q", user.Name, "Alice")
	}
	if user.Email != "alice@example.com" {
		t.Errorf("Email: got %q, want %q", user.Email, "alice@example.com")
	}
}

func TestDisplayName(t *testing.T) {
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	display := user.DisplayName()

	expected := "Alice <alice@example.com>"
	if display != expected {
		t.Errorf("got %q, want %q", display, expected)
	}
}

func TestUpdateEmail(t *testing.T) {
	user := &User{ID: 1, Name: "Alice", Email: "old@example.com"}
	user.UpdateEmail("new@example.com")

	if user.Email != "new@example.com" {
		t.Errorf("Email after update: got %q, want %q", user.Email, "new@example.com")
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"alice@example.com", true},
		{"test@test", true},
		{"invalid", false},
		{"", false},
		{"has@symbol", true},
	}

	for _, tc := range tests {
		user := User{Email: tc.email}
		result := user.IsValidEmail()
		if result != tc.expected {
			t.Errorf("IsValidEmail(%q): got %v, want %v", tc.email, result, tc.expected)
		}
	}
}

func TestNewAdmin(t *testing.T) {
	admin := NewAdmin(1, "Bob", "bob@example.com", "superadmin")

	if admin == nil {
		t.Fatal("NewAdmin returned nil")
	}

	// Check embedded User fields
	if admin.ID != 1 {
		t.Errorf("ID: got %d, want 1", admin.ID)
	}
	if admin.Name != "Bob" {
		t.Errorf("Name: got %q, want %q", admin.Name, "Bob")
	}
	if admin.Role != "superadmin" {
		t.Errorf("Role: got %q, want %q", admin.Role, "superadmin")
	}
}

func TestAdminInheritsUserMethods(t *testing.T) {
	admin := Admin{
		User: User{ID: 1, Name: "Bob", Email: "bob@example.com"},
		Role: "admin",
	}

	// Admin should have DisplayName from User
	display := admin.DisplayName()
	expected := "Bob <bob@example.com>"
	if display != expected {
		t.Errorf("Admin.DisplayName(): got %q, want %q", display, expected)
	}
}

func TestCanDelete(t *testing.T) {
	superadmin := Admin{User: User{}, Role: "superadmin"}
	if !superadmin.CanDelete() {
		t.Error("superadmin should be able to delete")
	}

	regularAdmin := Admin{User: User{}, Role: "admin"}
	if regularAdmin.CanDelete() {
		t.Error("regular admin should not be able to delete")
	}
}

func TestNewProduct(t *testing.T) {
	product := NewProduct(1, "Widget", 29.99)

	if product.ID != 1 {
		t.Errorf("ID: got %d, want 1", product.ID)
	}
	if product.Name != "Widget" {
		t.Errorf("Name: got %q, want %q", product.Name, "Widget")
	}
	if product.Price != 29.99 {
		t.Errorf("Price: got %f, want %f", product.Price, 29.99)
	}
}

func TestWithDiscount(t *testing.T) {
	original := Product{ID: 1, Name: "Widget", Price: 100.0}
	discounted := original.WithDiscount(20)

	// Original should be unchanged
	if original.Price != 100.0 {
		t.Errorf("Original price changed: got %f, want 100.0", original.Price)
	}

	// Discounted should have new price
	if discounted.Price != 80.0 {
		t.Errorf("Discounted price: got %f, want 80.0", discounted.Price)
	}

	// Other fields should be same
	if discounted.ID != original.ID || discounted.Name != original.Name {
		t.Error("Discounted product should keep same ID and Name")
	}
}

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()

	if area != 50 {
		t.Errorf("Area: got %f, want 50", area)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 5}
	perimeter := rect.Perimeter()

	if math.Abs(perimeter-30) > 0.001 {
		t.Errorf("Perimeter: got %f, want 30", perimeter)
	}
}
