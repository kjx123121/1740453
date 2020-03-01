package calculation

func FindBodyShape(bmi float64, bodyShape map[string][2]float64) string {
	var shape string

	for k, v := range bodyShape {

		if bmi > v[0] && bmi < v[1] {
			shape = k
		}
	}

	return shape
}