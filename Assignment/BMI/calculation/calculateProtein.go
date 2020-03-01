package calculation

import "math"

func CalculateProtein(diet bool, weight float64) (float64, float64) {
	var minimumProtein, maximumProtein float64

	if diet {
		minimumProtein = weight * 1.2
		maximumProtein = weight * 1.5

	} else {
		minimumProtein = weight * 1
		maximumProtein = weight * 1.2
	}

	return math.Round(minimumProtein*100) / 100, math.Round(maximumProtein*100) / 100
}