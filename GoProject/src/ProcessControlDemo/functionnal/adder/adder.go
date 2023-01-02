package main

import "fmt"

func adder() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(v int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		sum := base + v
		return sum, adder2(sum)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	fmt.Println("-----------------------")

	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
