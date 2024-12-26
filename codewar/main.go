package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := NumberToString(54546)
	fmt.Println(s)
}

func NumberToString(n int) string {
	return strconv.Itoa(n)
}
