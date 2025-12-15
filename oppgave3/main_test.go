package main

import "testing"

func TestIsAPrime(t *testing.T) {
	result := IsAPrime(2)

	expected := true

	if result != expected {
		t.Errorf("IsAPrime(2) = %v; vil ha %v", result, expected)

	}

}
