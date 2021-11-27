package main

import (
	"fmt"
	"math"
)

func baito_kyuyo (year int) int {
	base := 850
	year_bonus := year * 100
	return base + year_bonus
}

func jikoshokai (name string) string {
	return fmt.Sprintf("こんにちは、%s", name)
}

func hyojun_taiju (height float64) float64 {
	return height * height * 22
}

func bmi(height, weight float64) float64 {
	bmi := weight / (height * height)
	bmiround := (math.Round(bmi * 100) / 100)
	return bmiround
}

func main() {
	fmt.Println(baito_kyuyo(10))
	fmt.Println(jikoshokai("和田"))
	fmt.Println(hyojun_taiju(1.66))
	fmt.Println(bmi(1.66, 60))
}