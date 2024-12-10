package main

import (
	"fmt"
)

func main() {
	fmt.Println("apbqcr" == mergeAlternately("abc", "pqr"), mergeAlternately("abc", "pqr"))
	fmt.Println("apbqrs" == mergeAlternately("ab", "pqrs"), mergeAlternately("ab", "pqrs"))
	fmt.Println("apbqcd" == mergeAlternately("abcd", "pq"), mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1, word2 string) string {
	result := ""
	i, j := 0, 0

	for i < len(word1) || j < len(word2) {

		if i < len(word1) {
			result += string(word1[i])
			i++
		}

		if j < len(word2) {
			result += string(word2[j])
			j++
		}
	}

	return result
}
