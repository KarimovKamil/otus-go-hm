package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordInfo struct {
	Word      string
	Frequency int
}

func Top10(text string) []string {
	wordFreq := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		wordFreq[word]++
	}

	wordFreqSlice := make([]WordInfo, 0, len(wordFreq))
	for k, v := range wordFreq {
		wordFreqSlice = append(wordFreqSlice, WordInfo{k, v})
	}

	sort.Slice(wordFreqSlice, func(i, j int) bool {
		return compareWords(wordFreqSlice[i], wordFreqSlice[j])
	})

	result := make([]string, 0)
	for i := 0; i < 10 && i < len(wordFreqSlice); i++ {
		result = append(result, wordFreqSlice[i].Word)
	}
	return result
}

func compareWords(word1, word2 WordInfo) bool {
	if word1.Frequency == word2.Frequency {
		return strings.Compare(word1.Word, word2.Word) < 0
	}
	return word1.Frequency > word2.Frequency
}
