package main

import "strings"

func countDistinctWords(messages []string) int {
	wordSet := make(map[string]struct{})
	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			wordSet[lowerWord] = struct{}{}
		}
	}
	return len(wordSet)
}
