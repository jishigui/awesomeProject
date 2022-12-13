package main

import "fmt"

func main() {

	numbers := [...]int{1, 2, 3, 4}
	maxi := -1
	maxValue := -1
	for i, v := range numbers {
		if v > maxValue {
			maxi, maxValue = i, v
			fmt.Println(i, v)
		}
	}
	fmt.Println(maxi, maxValue)
}
