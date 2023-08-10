package main

import "fmt"

func Mapping(slice []string) map[string]int {
	count := make(map[string]int)
	for _, val := range slice {
		count[val]++
	}
	return count
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{}))
}
