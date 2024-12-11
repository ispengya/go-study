package main

import "fmt"

func arrayDemo() {
	var numArray = [...]int{1, 2, 5: 3}
	strArray := [...]string{0: "han", 1: "peng"}

	for i := 0; i < len(numArray); i++ {
		fmt.Println(numArray[i])
	}

	for index, value := range strArray {
		fmt.Println(index, value)
	}

	a := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i, val := range a {
		fmt.Println(i, val)
	}
}

func main() {
	arrayDemo()
}
