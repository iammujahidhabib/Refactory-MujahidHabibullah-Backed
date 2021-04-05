package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getCountWord() {
	var paragraph = `Aku ingin begini
	Aku ingin begitu
	Ingin ini itu banyak sekali
	
	Semua semua semua
	Dapat dikabulkan
	Dapat dikabulkan
	Dengan kantong ajaib
	
	Aku ingin terbang bebas
	Di angkasa
	Hei… baling baling bambu
	
	La... la... la...
	Aku sayang sekali
	Doraemon
	
	La... la... la...
	Aku sayang sekali`

	str := [3]string{"aku", "ingin", "dapat"}

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i] + " : " + strconv.Itoa(strings.Count(strings.ToLower(paragraph), str[i])))
	}
}
func main() {
	fmt.Println("3. Count words")
	getCountWord()
}
