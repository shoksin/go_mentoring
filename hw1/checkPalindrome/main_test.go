package main

import (
	"testing"
)

func TestSyncCheckPalindrome(t *testing.T) {
	tests := map[string]bool{
		"":                              true,
		"hello":                         false,
		"aba":                           true,
		"xyyx":                          true,
		"uu":                            true,
		"a":                             true,
		"782ddd287":                     true,
		"m00010m":                       false,
		"abcdefghijklmnonmlkjihgfedcba": true,
	}
	for word, expected := range tests {
		if result := isPalindromeSimpleCheck(word); result != expected {
			t.Errorf("isPolindromeSimpleCheck(%q) = %v, want = %v", word, result, expected)
		}
	}
}

func TestAsyncCheckPalindrome(t *testing.T) {
	tests := map[string]bool{
		"":                              true,
		"hello":                         false,
		"aba":                           true,
		"xyyx":                          true,
		"uu":                            true,
		"a":                             true,
		"782ddd287":                     true,
		"m00010m":                       false,
		"abcdefghijklmnonmlkjihgfedcba": true,
	}
	for word, expected := range tests {
		if result := isPalindromeAsyncCheck(word); result != expected {
			t.Errorf("isPolindromeSimpleCheck(%q) = %v, want = %v", word, result, expected)
		}
	}
}
