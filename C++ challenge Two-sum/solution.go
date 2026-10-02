package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		needed := target - num

		if index, exists := seen[needed]; exists {
			return []int{index, i}
		}

		seen[num] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println(result)
}
