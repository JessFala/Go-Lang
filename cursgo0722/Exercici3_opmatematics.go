package main

import (
	"fmt"
)

var (
	c int     = 83
	j int     = 13
	l float32 = 16.0
	t float32 = 12.5
)

func main() {
	fmt.Println(c + j)
	fmt.Println(j - c)
	fmt.Println(l * t)
	fmt.Println(l / t)
}
