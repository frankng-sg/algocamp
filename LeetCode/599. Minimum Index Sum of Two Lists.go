package main

import "fmt"

func findRestaurant(list1, list2 []string) []string {
	index := make(map[string]int)
	for i, v := range list1 {
		index[v] = i
	}
	min := len(list1) + len(list2)
	var result []string
	for i, v := range list2 {
		if j, ok := index[v]; ok {
			if i+j < min {
				min = i + j
				result = []string{v}
			} else if i+j == min {
				result = append(result, v)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(findRestaurant(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
	)) // Output: Shogun

	fmt.Println(findRestaurant(
		[]string{"happy", "sad", "good"},
		[]string{"sad", "happy", "good"},
	)) // Output: sad, happy
}
