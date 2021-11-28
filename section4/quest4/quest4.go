package quest4

import (
	"fmt"
	"math"
)


func Baito_kyuyo (year int) int {
	base := 850
	year_bonus := year * 100
	return base + year_bonus
}

func Jikoshokai (name string) string {
	return fmt.Sprintf("こんにちは、%s", name)
}

func Hyojun_taiju (height float64) float64 {
	return height * height * 22
}

func Bmi(height, weight float64) float64 {
	bmi := weight / (height * height)
	bmiround := (math.Round(bmi * 100) / 100)
	return bmiround
}