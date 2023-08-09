package main

import "fmt"

func main() {
	var atas, bawah, tinggi float64

	fmt.Print("Masukkan panjang sisi atas trapesium: ")
	fmt.Scan(&atas)

	fmt.Print("Masukkan panjang sisi bawah trapesium: ")
	fmt.Scan(&bawah)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)

	luas := (atas + bawah) * tinggi / 2

	fmt.Printf("Luas trapesium: %.2f\n", luas)
}
