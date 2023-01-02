package main

import (
	"awesomeProject/src/ProcessControlDemo/retriever/mock"
	real2 "awesomeProject/src/ProcessControlDemo/retriever/real"
	"fmt"
	"time"
)

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retiever
	r = mock.Retriever{Contens: "this is a fake imooc.com"}
	//fmt.Println(download(r))
	inspect(r)
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contens)
	} else {
		fmt.Println("not a mock retriever")
	}

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(download(r))
	inspect(r)

	// Type assertion
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)
}

func inspect(r Retiever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contens)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
