package main

import (
	"fmt"
)

type numList []int

func main() {
	list := make([]int, 11)

	for i := range list {
		list[i] = i
	}

	for num := range list {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
