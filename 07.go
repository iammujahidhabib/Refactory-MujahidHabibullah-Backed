package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{9, 4, 2, 4, 1, 5, 3, 0}
	new_numbers := []int{}
	fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 != 0 {
			new_numbers = append(new_numbers, numbers[i])
		}
	}
	// fmt.Println(new_numbers)
	a := 0
	// fmt.Println(new_numbers[a])
	sort.Slice(new_numbers, func(i, j int) bool {
		return new_numbers[i] < new_numbers[j]
	})
	fmt.Println(new_numbers)
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 != 0 {
			numbers[i] = new_numbers[a]
			a++
		}
	}
	fmt.Println(numbers)
}
