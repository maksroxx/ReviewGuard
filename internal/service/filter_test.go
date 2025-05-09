package service

import (
	"testing"
)

func TestContainsBannedWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Clean sentence", "Это лучший сервис в мире", false},
		{"Exact match", "нigger", false},
		{"Mixed case", "NiGgEr", true},
		{"Embedded in word", "helloNIGGERworld", true},
		{"Unicode + dirty symbols", "ГаЙ...", false},
		{"Censored word", "г а й", false},
		{"Multiple words", "Это глупо и дурак", true},
		{"Partial root", "дура", false},
		{"Trailing punctuation", "gay!", true},
		{"Clean mixed", "ничего плохого тут нет", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsBannedWords(tc.input)
			if result != tc.expected {
				t.Errorf("ContainsBannedWords(%q) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestContainsLinks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"HTTP link", "Check http://example.com", true},
		{"HTTPS link", "Go to https://secure.org now", true},
		{"No link", "Just some random words", false},
		{"Invalid URL", "visit htps:/wrong.com", false},
		{"Embedded link", "text before https://google.com and after", true},
		{"Link with punctuation", "https://example.com!", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsLinks(tc.input)
			if result != tc.expected {
				t.Errorf("ContainsLinks(%q) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
