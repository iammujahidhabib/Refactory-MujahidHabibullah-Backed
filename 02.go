package main

import (
	"fmt"
	"strings"
)

func CharacterCount1(s string) int {
	return len(s)
}
func getStars(s int) string {
	var sss = ""
	for i := 0; i < s; i++ {
		sss += "*"
	}
	return sss
}
func getCencoring() {
	var paragraph = `Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
	Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
	Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	str := [7]string{"dolor", "elit", "quis", "nisi", "fugiat", "proident", "laborum"}

	for i := 0; i < len(str); i++ {
		paragraph = strings.Replace(paragraph, str[i], getStars(CharacterCount1(str[i])), 1)
	}
	fmt.Println(paragraph)
}
func main() {
	fmt.Println("2. Censoring Words")
	getCencoring()
}
