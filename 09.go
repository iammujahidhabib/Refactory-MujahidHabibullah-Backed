package main

import (
	"fmt"
	"sort"
)

func main() {
	arrayList := []int{3, 12, 4, 5, 8, 9}
	sort.Slice(arrayList, func(i, j int) bool {
		return arrayList[i] < arrayList[j]
	})

	fmt.Println("9. Array Sort")
	fmt.Println(arrayList)
}
