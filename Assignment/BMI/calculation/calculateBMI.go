package calculation

import "math"

func CalculateBMI(height, weight float64) float64 {
	bmi := weight / math.Pow(height/100, 2)
	return math.Round(bmi*10) / 10
}