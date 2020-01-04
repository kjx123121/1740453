package main

import "fmt"

func main() {
	var originStr = "aabccdaa"

	var returnSlice []byte
	var counter uint8 = '1'
	for k, _ := range originStr {
		if k != len(originStr) - 1 {
			if originStr[k] == originStr[k + 1] {
				counter += 1
			} else {
				returnSlice = append(returnSlice, originStr[k], counter)
				counter = '1'
			}
		} else {
			returnSlice = append(returnSlice, originStr[k], counter)
		}

	}
	fmt.Printf("%s\n", returnSlice)
}
