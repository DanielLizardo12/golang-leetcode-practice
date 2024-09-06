package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for index, val := range nums {
		if _, ok := seen[target-val]; ok {
			return []int{seen[target-val], index}
		}
		seen[val] = index
	}

	return nil
}
