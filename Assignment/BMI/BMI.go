package main

import (
	"Assignment/BMI/calculation"
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


	height := getHeight()
	weight := getWeight()
	exercise := getExercise()
	diet := getDiet()



	bmi := calculation.CalculateBMI(height, weight)
	fmt.Printf("您的BMI为: %0.1f\n", bmi)

	shape := calculation.FindBodyShape(bmi, bodyShape)
	fmt.Printf("属于%s体型\n", shape)

	minEnergy, maxEnergy := calculation.CalculateEnergy(weight, adultEnergy, shape, exercise)
	fmt.Printf("每日摄取热量范围为: %0.2f--%0.2f(kJ)\n", minEnergy, maxEnergy)

	minProtein, maxProtein := calculation.CalculateProtein(diet, weight)
	fmt.Printf("每日摄取蛋白质范围为: %0.2f--%0.2f(g)\n", minProtein, maxProtein)

}

func stdinInt() int {
	var intReceiver int
	_, err := fmt.Scanf("%d", &intReceiver)

	if err != nil {
		fmt.Println(err)

	}
	return intReceiver
}

func stdinFloat() float64 {
	var floatReceiver float64

	_, err := fmt.Scanf("%f", &floatReceiver)

	if err != nil {
		fmt.Println(err)

	}
	return floatReceiver
}

func stdinBoolean() bool {
	var booleanReceiver bool

	_, err := fmt.Scanf("%t", &booleanReceiver)

	if err != nil {
		fmt.Println(err)

	}
	return booleanReceiver
}


func getHeight() float64{
	var height float64
	for {
		fmt.Print("请输入身高(cm) : ")
		height = stdinFloat()
		if height > 0 {
			break
		}else {
			fmt.Printf("请输入正确的身高\n")
		}
	}
	return height
}
func getWeight() float64{
	var weight float64
	for {
		fmt.Print("请输入体重(kg) : ")
		weight = stdinFloat()
		if weight > 0 {
			break
		}else {
			fmt.Printf("请输入正确的体重\n")
		}
	}
	return weight
}
func getExercise() int{
	var exercise int
	for {
		fmt.Print("活动强度用 1 - 4 来表示 (轻体力劳动 -- 重体力劳动) : ")
		exercise = stdinInt()
		if exercise == 1 || exercise == 2 || exercise == 3 || exercise == 4{
			break
		}else {
			fmt.Printf("请仅使用1 - 4 来表示\n")
		}
	}
	return exercise
}
func getDiet() bool{
	var diet bool
	for {
		fmt.Print("是否断用主食(t:是 f:否) : ")
		diet = stdinBoolean()
		if diet == true || diet == false {
			break
		}else {
			fmt.Printf("请使用 't' 或 'f' 来回答\n")
		}
	}
	return diet
}