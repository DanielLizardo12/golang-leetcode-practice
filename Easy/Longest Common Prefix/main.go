package main

import (
	"fmt"
)

func main() {

	testArray := [3]string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(testArray[:]))
}

func longestCommonPrefix(strs []string) string {
	longestPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(string(strs[i])); j++ {
			fmt.Println(strs)
		}
	}

	return longestPrefix
}
