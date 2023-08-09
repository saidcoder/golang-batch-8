package main

import "fmt"

func main() {
	var num int
	var flag bool

	fmt.Print("Masukkan Bilangan: ")

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Input harus angka.")
		return
	}

	if num%7 == 0 {
		flag = true
	}

	if flag == true {
		fmt.Printf("%d bilangan dari kelipatan 7.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan dari kelipatan 7.\n", num)
	}
}
