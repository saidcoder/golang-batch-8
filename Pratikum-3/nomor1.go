package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make(map[string]bool)
	result := []string{}

	// Tambahkan elemen dari arrayA ke hasil
	for _, val := range arrayA {
		merged[val] = true
		result = append(result, val)
	}

	// Tambahkan elemen dari arrayB yang belum ada di hasil
	for _, val := range arrayB {
		if !merged[val] {
			merged[val] = true
			result = append(result, val)
		}
	}

	return result
}

func main() {
	//test casec
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
}
