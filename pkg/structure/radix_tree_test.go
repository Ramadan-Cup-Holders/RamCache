package structure

import (
	"testing"
)

func TestRadixTree(t *testing.T) {
	radix := NewRadixTree()

	// Test inserting words
	radix.Insert("hello", "world")
	radix.Insert("helium", "gas")
	radix.Insert("hero", "brave")
	radix.Insert("her", "she")

	// Test searching existing words
	value, found := radix.Search("hello")
	if !found || value != "world" {
		t.Errorf("Expected 'world', got '%s'", value)
	}

	value, found = radix.Search("helium")
	if !found || value != "gas" {
		t.Errorf("Expected 'gas', got '%s'", value)
	}

	value, found = radix.Search("hero")
	if !found || value != "brave" {
		t.Errorf("Expected 'brave', got '%s'", value)
	}

	value, found = radix.Search("her")
	if !found || value != "she" {
		t.Errorf("Expected 'she', got '%s'", value)
	}

	// Test searching non-existing word
	_, found = radix.Search("hex")
	if found {
		t.Errorf("Expected 'hex' to not be found")
	}

	// Test searching prefix that is not a complete word
	_, found = radix.Search("he")
	if found {
		t.Errorf("Expected 'he' to not be found")
	}
}
