package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeat(names ...string) {
	fmt.Println(names)
}
func main() {
	name := "mgkim"
	fmt.Println(name)

	len, upperName := lenAndUpper(name)
	fmt.Println(len, upperName)

	// 다중 인자 정의
	repeat("mgkim", "selee", "korea")
}
