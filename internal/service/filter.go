package service

import (
	"regexp"
	"strings"
)

var banWords = []string{"nigger", "gay", "дурак"}

var (
	urlRegex   = regexp.MustCompile(`https?://[^\s]+`)
	cleanRegex = regexp.MustCompile(`[^\p{L}]`)
)

func ContainsBannedWords(text string) bool {
	clean := cleanRegex.ReplaceAllString(strings.ToLower(text), "")
	for _, word := range banWords {
		if strings.Contains(clean, word) {
			return true
		}
	}
	return false
}

func ContainsLinks(text string) bool {
	return urlRegex.MatchString(text)
}
