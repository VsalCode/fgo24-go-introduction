package utils

import "fmt"

func hitungKeliling(r int) string {
	if r%7 == 0 {
		calculate := 2 * 22/7 * r
		return fmt.Sprintf("Hasil Luas Lingkaran : %d", calculate)
	}
	phi := 3.14
	result := 2 * phi * float64(r) 
	return fmt.Sprintf("Hasil Luas Lingkaran : %g", result)
}

func ResultKeliling() any {
	calc := hitungKeliling(7)
	return calc	
}
