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
func getCharCounter(s string) {
	charArray := []string{}
	// newCharArray := []string{}
	for _, c := range s {
		_, found := Find(charArray, string(c))
		if !found {
			charArray = append(charArray, string(c))
		}
	}
	fmt.Println(charArray)
	var m = make(map[string]string)
	// arr := []string{}
	for i := 0; i < len(charArray); i++ {
		m[charArray[i]] = getStars(strings.Count(s, charArray[i]))
	}
	fmt.Println(m)
}
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func main() {
	fmt.Println("5. Char Counter")

	var paragraph = `Mammals`
	var paragraph2 = `Bruiser build`
	getCharCounter(strings.ToLower(paragraph))
	getCharCounter(strings.ToLower(paragraph2))
}
