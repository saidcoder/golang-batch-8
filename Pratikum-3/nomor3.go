package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	countMap := make(map[int]int)

	for _, char := range angka {
		digit := int(char - '0')
		countMap[digit]++
	}

	munculSekali := []int{}

	for digit, count := range countMap {
		if count == 1 {
			munculSekali = append(munculSekali, digit)
		}
	}

	return munculSekali
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("11223344"))
}
