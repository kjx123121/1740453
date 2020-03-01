package getInput

import "fmt"

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

func GetHeight() float64 {
	var height float64

	for {

		fmt.Print("请输入身高(cm) : ")
		height = stdinFloat()

		if height > 0 {
			break

		} else {
			fmt.Printf("请输入正确的身高\n")
		}
	}
	return height
}
func GetWeight() float64 {
	var weight float64

	for {
		fmt.Print("请输入体重(kg) : ")
		weight = stdinFloat()

		if weight > 0 {
			break

		} else {
			fmt.Printf("请输入正确的体重\n")
		}
	}
	return weight
}
func GetExercise() int {
	var exercise int

	for {
		fmt.Print("活动强度用 1 - 4 来表示 (轻体力劳动 -- 重体力劳动) : ")
		exercise = stdinInt()

		if exercise == 1 || exercise == 2 || exercise == 3 || exercise == 4 {
			break

		} else {
			fmt.Printf("请仅使用1 - 4 来表示\n")
		}
	}
	return exercise
}
func GetDiet() bool {
	var diet bool

	for {
		fmt.Print("是否断用主食(t:是 f:否) : ")
		diet = stdinBoolean()

		if diet == true || diet == false {
			break

		} else {
			fmt.Printf("请使用 't' 或 'f' 来回答\n")
		}
	}
	return diet
}
