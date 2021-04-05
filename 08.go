package main

import (
	"fmt"
	"strings"
)

func getStars(s int) string {
	var sss = ""
	for i := 0; i < s; i++ {
		sss += "*"
	}
	return sss
}
func getMasking(s string) string {
	str := []rune(s)
	for i := 0; i < len(str)-3; i++ {
		str[i] = '*'
	}
	return string(str)
}
func main() {
	fmt.Println("8. Masking")

	var paragraph = `23dn3ir30fd2eddd`

	fmt.Println(getMasking(strings.ToLower(paragraph)))
}
