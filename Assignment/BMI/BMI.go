package main

import (
	"Assignment/BMI/calculation"
	"Assignment/BMI/getInput"
	"fmt"
	"math"
)

var bodyShape map[string][2]float64 = map[string][2]float64{
	"消瘦":    {0.0, 18.4},
	"正常":    {18.5, 23.9},
	"超重&肥胖": {24.0, math.Inf(1)},
}

var adultEnergy map[int]map[string][2]float64 = map[int]map[string][2]float64{
	1: {"消瘦": {25, 30}, "正常": {20, 25}, "超重&肥胖": {15, 20}},
	2: {"消瘦": {30, 35}, "正常": {25, 30}, "超重&肥胖": {20, 25}},
	3: {"消瘦": {35, 40}, "正常": {30, 35}, "超重&肥胖": {25, 30}},
	4: {"消瘦": {40, 45}, "正常": {35, 40}, "超重&肥胖": {30, 35}},
}

func main() {


	height := getInput.GetHeight()
	weight := getInput.GetWeight()
	exercise := getInput.GetExercise()
	diet := getInput.GetDiet()



	bmi := calculation.CalculateBMI(height, weight)
	fmt.Printf("您的BMI为: %0.1f\n", bmi)

	shape := calculation.FindBodyShape(bmi, bodyShape)
	fmt.Printf("属于%s体型\n", shape)

	minEnergy, maxEnergy := calculation.CalculateEnergy(weight, adultEnergy, shape, exercise)
	fmt.Printf("每日摄取热量范围为: %0.2f--%0.2f(kJ)\n", minEnergy, maxEnergy)

	minProtein, maxProtein := calculation.CalculateProtein(diet, weight)
	fmt.Printf("每日摄取蛋白质范围为: %0.2f--%0.2f(g)\n", minProtein, maxProtein)

}
