package main

import "fmt"

func main() {
	fmt.Println(hitungKeliling(7))
	fmt.Println(hitungLuas(7))
}

func hitungLuas(r int) string {
	if r%7 == 0 {
		calculate := 22 / 7 * r * r
		return fmt.Sprintf("Hasil Luas Lingkaran : %d", calculate)
	}
	phi := 3.14
	result := phi * float64(r) * float64(r)
	return fmt.Sprintf("Hasil Luas Lingkaran : %g", result)
}

func hitungKeliling(r int) string {
	if r%7 == 0 {
		calculate := 2 * 22/7 * r
		return fmt.Sprintf("Hasil Luas Lingkaran : %d", calculate)
	}
	phi := 3.14
	result := 2 * phi * float64(r) 
	return fmt.Sprintf("Hasil Luas Lingkaran : %g", result)
}
