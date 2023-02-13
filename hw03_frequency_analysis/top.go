package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordFrequency struct {
	word  string
	count int
}

func newWordFrequency(word string, count int) wordFrequency {
	return wordFrequency{
		word:  word,
		count: count,
	}
}

func Top10(text string) []string {
	words := strings.Fields(text)
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

	frequencies := make([]wordFrequency, 0, len(wordsMap))
	for key, val := range wordsMap {
		frequencies = append(frequencies, newWordFrequency(key, val))
	}

	sort.Slice(frequencies, func(i, j int) bool {
		a := frequencies[i]
		b := frequencies[j]
		if a.count == b.count {
			return a.word < b.word
		}
		return a.count > b.count
	})

	outSize := 10
	if outSize > len(frequencies) {
		outSize = len(frequencies)
	}

	top := make([]string, outSize)
	for i := 0; i < outSize; i++ {
		top[i] = frequencies[i].word
	}
	return top
}
