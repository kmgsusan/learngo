package main

import "fmt"

func main() {
	ss := make([]string, 3, 3)
	fmt.Println(ss)

	var a []int
	a = []int{1, 2, 3, 4, 5}
	a[0] = 0
	fmt.Println(a)

	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	x := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(x)
	x = s[2:5]
	fmt.Println(x)

	y := make([]int, 2, 3)
	fmt.Println(y)

}
