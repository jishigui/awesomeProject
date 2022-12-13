package main

import "fmt"

func main() {
	var s []int

	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	fmt.Println(s1)

	s2 := make([]int, 16)

	s3 := make([]int, 10, 32)

	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	fmt.Println(s2)

	fmt.Println("Deleting elements form slice")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
}
