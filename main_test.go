package main

import "testing"

// TestIsEven menguji fungsi IsEven
func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int  // Nilai masukan
		expected bool // Hasil yang diharapkan
	}{
		{1, false},
		{2, true},
		{0, true},
		{-1, false},
		{-2, true},
	}

	for _, tt := range tests {
		result := IsEven(tt.input)
		if result != tt.expected {
			t.Errorf("IsEven(%d) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
