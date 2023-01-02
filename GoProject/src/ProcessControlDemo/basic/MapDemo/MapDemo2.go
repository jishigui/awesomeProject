package main

import "fmt"

func main() {
	rating := map[string]float32{"c": 5, "Go": 4.5, "python": 4.5, "c++": 2}

	csharpRating, ok := rating["c#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")

}
