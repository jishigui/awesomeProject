package main

import "fmt"

/*
*
slice 的默认开始位置是 0，ar[:n] 等价于 ar[0:n]
slice 的第二个序列默认是数组的长度，ar[n:] 等价于 ar[n:len(ar)]
如果从一个数组里面直接获取 slice，可以这样 ar[:]，因为默认第一个序列是 0，第二个是数组的长度，即等价于 ar[0:len(ar)]
*/
func main() {
	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个 slice
	var aSlice, bSlice, cSlice []byte

	aSlice = array[:3]
	fmt.Println(aSlice)
	aSlice = array[5:]
	fmt.Println(aSlice)
	aSlice = array[:]
	fmt.Println(aSlice)

	aSlice = array[3:7]
	fmt.Println(aSlice)
	bSlice = aSlice[1:3]
	fmt.Println(bSlice)
	bSlice = aSlice[:3]
	fmt.Println(bSlice)

	bSlice = aSlice[0:5]
	fmt.Println(bSlice)
	aSlice[0] = '4'
	fmt.Println("---", bSlice)

	bSlice = aSlice[:]
	fmt.Println(bSlice)

	fmt.Println(aSlice, bSlice)

	i := copy(aSlice, cSlice)
	fmt.Println(i)
}
