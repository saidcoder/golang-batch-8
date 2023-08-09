package main

import (
	"fmt"
)

func main() {
	var num int
	var flag bool

	fmt.Print("Masukkan Bilangan: ")

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input harus angka.")
		return
	}

	for count := 2; count <= num/2; count++ {
		if num%count == 0 {
			flag = true
			break
		}
	}

	if flag == true || num < 2 {
		fmt.Printf("%d bukan bilangan prima.\n", num)
	} else {
		fmt.Printf("%d adalah bilangan prima.\n", num)
	}

}
