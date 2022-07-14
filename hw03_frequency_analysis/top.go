package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordCountStruct struct {
	w string
	c int32
}

const TopNumber = 10

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Top10(inp string) []string {
	tokenFrequency := make(map[string]int32)
	r := regexp.MustCompile("[a-яА-Я-]+")
	tokens := strings.Fields(inp)
	for _, t := range tokens {
		if r.MatchString(t) {
			tokenFrequency[t]++
		}
	}
	wordCounts := []wordCountStruct{}
	for key, val := range tokenFrequency {
		wordCounts = append(wordCounts, wordCountStruct{w: key, c: val})
	}

	maxRetLen := min(TopNumber, len(wordCounts))

	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].c == wordCounts[j].c {
			return wordCounts[i].w < wordCounts[j].w
		}
		return wordCounts[i].c > wordCounts[j].c
	})
	result := make([]string, maxRetLen)
	for i, curWordCount := range wordCounts[:maxRetLen] {
		result[i] = curWordCount.w
	}
	return result
}
