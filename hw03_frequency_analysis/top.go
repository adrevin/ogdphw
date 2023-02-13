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

func Top10(inp string) []string {

	words := strings.Fields(inp)
	var wordsMap = make(map[string]int)
	for _, word := range words {
		if _, ok := wordsMap[word]; ok {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
		}
	}

	var frequencies = make([]wordFrequency, len(wordsMap))
	i := 0
	for key, val := range wordsMap {
		frequencies[i] = newWordFrequency(key, val)
		i++
	}

	sort.Slice(frequencies, func(i, j int) bool {
		a := frequencies[i]
		b := frequencies[j]
		if a.count == b.count {
			return a.word < b.word
		}
		return a.count > b.count
	})

	var outSize = 10
	var frequenciesCount = len(frequencies)
	if outSize > frequenciesCount {
		outSize = frequenciesCount
	}

	out := make([]string, outSize)
	for i := 0; i < outSize; i++ {
		out[i] = frequencies[i].word
	}

	return out
}
