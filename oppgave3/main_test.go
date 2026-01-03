package main

import "testing"

func TestIsAPrime(t *testing.T) {
	// Test at disse tallene er primtall
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	
	for _, prime := range primes {
		if !IsAPrime(prime) {
			t.Errorf("IsAPrime(%d) = false; vil ha true", prime)
		}
	}
}

