package main

import (
	"fmt"
)

func getEvenOdd(s []int) {
	// fmt.Println(len(s))
	intEven := []int{}
	intOdd := []int{}
	intByFive := []int{}
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			intEven = append(intEven, i)
		} else {
			intOdd = append(intOdd, i)
		}
		intByFive = append(intByFive, i*5)
	}
	fmt.Print("Even : ")
	fmt.Print(intEven)
	fmt.Println("")
	fmt.Print("Odd : ")
	fmt.Print(intOdd)
	fmt.Println("")
	fmt.Print("numbers multiplies by 5 : ")
	fmt.Print(intByFive)
	fmt.Println("")
}
func getPrime(s []int) {
	// fmt.Println(len(s))
	intPrime := []int{}
	intPrimeLess := []int{}
	for i := 0; i <= len(s); i++ {
		var a = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				a++
			}
		}
		if a == 2 {
			intPrime = append(intPrime, i)
			if i < 100 {
				intPrimeLess = append(intPrimeLess, i)
			}

		}
	}
	fmt.Print("prime numbers : ")
	fmt.Print(intPrime)
	fmt.Println("")
	fmt.Print("prime numbers less than 100. : ")
	fmt.Print(intPrimeLess)
	fmt.Println("")
}
func main() {
	fmt.Println("3. Count words")
	intArray := []int{}
	for i := 0; i <= 1000; i++ {
		intArray = append(intArray, i)
	}
	fmt.Print("row of numbers : ")
	fmt.Print(intArray)
	fmt.Println("")
	getEvenOdd(intArray)
	getPrime(intArray)
}
