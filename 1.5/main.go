package main

import "fmt"

func numbers(nums ...int) {
	for index, d := range nums {
		fmt.Println("%d - %d", index, d)
	}
}
func main() {
	numbers(1, 2, 3, 4, 5, 6, 7)
}
