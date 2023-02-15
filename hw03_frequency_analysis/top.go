package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)
	wordsCounts := make(map[string]int)

	for _, word := range words {
		wordsCounts[word]++
	}

	sortedWords := make([]string, 0, len(wordsCounts))

	for word := range wordsCounts {
		sortedWords = append(sortedWords, word)
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		a := sortedWords[i]
		b := sortedWords[j]
		if wordsCounts[a] == wordsCounts[b] {
			return a < b
		}
		return wordsCounts[a] > wordsCounts[b]
	})

	outSize := 10
	if outSize > len(sortedWords) {
		outSize = len(sortedWords)
	}
	return sortedWords[:outSize]
}
