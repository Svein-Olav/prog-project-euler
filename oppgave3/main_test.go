package main

import "testing"

func TestIsAPrime(t *testing.T) {
	/*
	result := IsAPrime(2)
	expected := true

	if result != expected {
		t.Errorf("IsAPrime(2) = %v; vil ha %v", result, expected)

	}
*/
	expected := false
	if IsAPrime(17) != expected {
		t.Errorf("IsAPrime(10) = %v; vil ha %v", IsAPrime(11), false)
	}

	/*
	expected = true
	if IsAPrime(1009) != true {
		t.Errorf("IsAPrime(1009) = %v; vil ha %v", IsAPrime(1009), true)
	}
		*/

}

func TestFinnAllePrimtallOpptil(t *testing.T) {
	result := FinnAllePrimtallOpptil(28)
	expected := []int{2, 7}

	if len(result) != len(expected) {
		t.Errorf("FinnAllePrimtallOpptil(28) = %v; vil ha %v", result, expected)
		return
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("FinnAllePrimtallOpptil(28) = %v; vil ha %v", result, expected)
			return
		}
	}

}