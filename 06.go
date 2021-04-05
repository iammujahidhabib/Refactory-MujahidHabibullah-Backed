package main

import (
	"fmt"
	"sort"
)

func MissingLetter(s []string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var first = ""
	var last = ""
	var finding string
	first = s[0]
	last = s[len(s)-1]
	idx := []int{}
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] == first || alphabet[i] == last {
			idx = append(idx, i)
		}
	}
	sort.Slice(idx, func(i, j int) bool {
		return idx[i] < idx[j]
	})
	// fmt.Println(idx)
	// fmt.Println(last)

	alphabet = alphabet[idx[0]:idx[1]]
	fmt.Println(alphabet)
	for j := 0; j < len(alphabet); j++ {
		if alphabet[j] != s[j] {
			finding = alphabet[j]
		}
	}
	return finding
}

func main() {

	list_letters_first := []string{"c", "d", "e", "g", "h"}
	list_letters_second := []string{"X", "Z"}

	fmt.Println(MissingLetter(list_letters_first))
	fmt.Println(MissingLetter(list_letters_second))
	// MissingLetter(list_letters_first)
}
