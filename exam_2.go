package main

import (
	"regexp"
	"strings"
)

func secondExamSolution(S string) int {
	sentencesArray := regexp.MustCompile("[\\.\\?\\!]").Split(S, -1)
	largestNumberOfWord := 0
	for _, sentence := range sentencesArray {
		currentNumberOfWord := len(strings.Split(strings.TrimSpace(sentence), " "))
		if largestNumberOfWord < currentNumberOfWord {
			largestNumberOfWord = currentNumberOfWord
		}
	}
	return largestNumberOfWord
}
