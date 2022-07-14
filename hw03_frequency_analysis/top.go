package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordCountStruct struct {
	w string // word
	c int32  // word count (frequency)
}

const TopNumber = 10

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Get token frequency as a map.
func getTokenFrequency(inp string, r *regexp.Regexp) map[string]int32 {
	tokenFrequency := make(map[string]int32)
	tokens := strings.Fields(inp)
	for _, t := range tokens {
		if r.MatchString(t) {
			tokenFrequency[t]++
		}
	}
	return tokenFrequency
}

// Get sorted wordCounts as a slice of wordCountStruct.
func getWordCounts(tokenFrequency map[string]int32) []wordCountStruct {
	// form wordCounts
	wordCounts := []wordCountStruct{}
	for key, val := range tokenFrequency {
		wordCounts = append(wordCounts, wordCountStruct{w: key, c: val})
	}

	// sort wordCounts
	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].c == wordCounts[j].c {
			return wordCounts[i].w < wordCounts[j].w
		}
		return wordCounts[i].c > wordCounts[j].c
	})

	return wordCounts
}

// Get topN words with max frequency.
func getTopWords(wordCounts []wordCountStruct, topN int) []string {
	result := make([]string, topN)
	for i, curWordCount := range wordCounts[:topN] {
		result[i] = curWordCount.w
	}
	return result
}

func Top10(inp string) []string {
	r := regexp.MustCompile("[a-яА-Я-]+")
	// get token frequency as map
	tokenFrequency := getTokenFrequency(inp, r)
	// get sorted wordCounts as a slice of wordCountStruct
	wordCounts := getWordCounts(tokenFrequency)
	// get max ret len
	maxRetLen := min(TopNumber, len(wordCounts))
	// get topN words with max frequency
	return getTopWords(wordCounts, maxRetLen)
}
