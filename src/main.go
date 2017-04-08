package main

import (
	"fmt"
)

var (
	name string
	age int
	location string
	add func(int, int) int
)

func main() {
	name, location = "Prince Oberyn", "Dorne"
	age = 32
	fmt.Printf("%s (%d) of %s\n", name, age, location)

	add = func(a int, b int) int {
		return a + b
	}
	fmt.Printf("1 + 2 = %d\n", add(1, 2))
}
