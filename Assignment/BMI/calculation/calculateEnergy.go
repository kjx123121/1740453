package calculation

import "math"

func CalculateEnergy(weight float64, adultEnergy map[int]map[string][2]float64, shape string, exercise int) (float64, float64) {
	var minimumEnergy, maximumEnergy float64

	minimumEnergy = weight * 0.7 * adultEnergy[exercise][shape][0]
	maximumEnergy = weight * 0.7 * adultEnergy[exercise][shape][1]

	return math.Round(minimumEnergy*100) / 100, math.Round(maximumEnergy*100) / 100
}