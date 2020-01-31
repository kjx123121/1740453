package main

import (
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

	fmt.Print("请输入身高(cm) : ")
	height := stdinFloat()

	fmt.Print("请输入体重(kg) : ")
	weight := stdinFloat()

	fmt.Print("活动强度用 1 - 4 来表示 (轻体力劳动 -- 重体力劳动) : ")
	exercise := stdinInt()

	fmt.Print("是否断用主食(1:是 2:否) : ")
	diet := stdinBoolean()

	bmi := calculateBMI(height, weight)
	fmt.Printf("您的BMI为: %0.1f\n", bmi)

	shape := findBodyShape(bmi, bodyShape)
	fmt.Printf("属于%s体型\n", shape)

	minEnergy, maxEnergy := calculateEnergy(weight, adultEnergy, shape, exercise)
	fmt.Printf("每日摄取热量范围为: %0.2f--%0.2f(kJ)\n", minEnergy, maxEnergy)

	minProtein, maxProtein := calculateProtein(diet, weight)
	fmt.Printf("每日摄取蛋白质范围为: %0.2f--%0.2f(g)\n", minProtein, maxProtein)

}

func stdinInt() int {
	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)

	}
	return i
}

func stdinFloat() float64 {
	var f float64

	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println(err)

	}
	return f
}

func stdinBoolean() bool {
	var b bool

	_, err := fmt.Scanf("%t", &b)

	if err != nil {
		fmt.Println(err)

	}
	return b
}

func calculateBMI(height, weight float64) float64 {
	bmi := weight / math.Pow(height/100, 2)
	return math.Round(bmi*10) / 10
}

func findBodyShape(bmi float64, bodyShape map[string][2]float64) string {
	var shape string

	for k, v := range bodyShape {

		if bmi > v[0] && bmi < v[1] {
			shape = k
		}
	}

	return shape
}

func calculateEnergy(weight float64, adultEnergy map[int]map[string][2]float64, shape string, exercise int) (float64, float64) {
	var minimumEnergy, maximumEnergy float64

	minimumEnergy = weight * 0.7 * adultEnergy[exercise][shape][0]
	maximumEnergy = weight * 0.7 * adultEnergy[exercise][shape][1]

	return math.Round(minimumEnergy*100) / 100, math.Round(maximumEnergy*100) / 100
}

func calculateProtein(diet bool, weight float64) (float64, float64) {
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
