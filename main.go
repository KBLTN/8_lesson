package main

import (
	tt "8_lesson/test"
	"fmt"
)

func main() {
	fmt.Println("test")
	tt.Add(1, 1)
}

func Multiple(x, y int) int {
	return x * y
}
