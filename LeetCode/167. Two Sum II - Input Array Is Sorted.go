package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // Output: [1, 2]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // Output: [1, 3]
	fmt.Println(twoSum([]int{-1, 0}, -1))       // Output: [1, 2]
}
