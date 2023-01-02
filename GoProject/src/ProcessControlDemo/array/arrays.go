package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	printArray(arr3)

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]       // [2,3,4,5]
	s2 := s1[2:6]        // [4,5,6,7]
	s3 := append(s2, 10) //
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(arr, s1, s2, s3, s4, s5)
}
