package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContentens(file)
}

func printFileContentens(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(72387885),
		convertToBin(0))

	printFile("abc.txt")

	s := `abc"d"
	kkkk
	123
	
	p`
	printFileContentens(strings.NewReader(s))
}
