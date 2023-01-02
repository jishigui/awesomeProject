package main

import "fmt"

func main() {

	slice := []byte{'a', 'b', 'c', 'd'}

	fmt.Println(slice)
	// 声明一个含有 10 个元素元素类型为 byte 的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// 声明两个含有 byte 的 slice
	var a, b []byte

	// a 指向数组的第 3 个元素开始，并到第五个元素结束，
	a = ar[2:5]
	// 现在 a 含有的元素: ar[2]、ar[3]和ar[4]

	// b 是数组 ar 的另一个 slice
	b = ar[3:5]
	// b 的元素是：ar[3] 和 ar[4]

	fmt.Println(a, b)
}
