package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, upperName string) {
	length = len(name)
	upperName = strings.ToUpper(name)
	return
}
func main() {
	name := "mgkim"
	lens, uname := lenAndUpper(name)
	fmt.Printf("%d, %s", lens, uname)
}
