package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello\nworld how are you?"
	index := strings.Index(str, " ")
	if index == -1 {
		index = strings.Index(str, "\n")
	}
	if index != -1 {
		str = str[0:index]
	}
	fmt.Println(str)
}
