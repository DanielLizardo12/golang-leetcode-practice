package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	romanValues := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	counter := 0

	for index := 0; index < len(s); index++ {
		if index+1 != len(s) {
			if romanValues[string(s[index+1])] > romanValues[string(s[index])] {
				counter = counter + (romanValues[string(s[index+1])] - romanValues[string(s[index])])
				index++
			} else {
				counter = counter + romanValues[string(s[index])]
			}
		} else {
			counter = counter + romanValues[string(s[index])]
		}
	}
	return counter
}
