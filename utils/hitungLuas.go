package utils

import "fmt"

func hitungLuas(r int) string {
	if r%7 == 0 {
		calculate := 22 / 7 * r * r
		return fmt.Sprintf("Hasil Luas Lingkaran : %d", calculate)
	}
	phi := 3.14
	result := phi * float64(r) * float64(r)
	return fmt.Sprintf("Hasil Luas Lingkaran : %g", result)
}

func ResultLuas() any {
	calc := hitungLuas(7)
	return calc	
}