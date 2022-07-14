package main

import (
	"fmt"
	"latohan-golang-part1/calculation"
)

func main() {
	fmt.Println("Halo, belajar golang")

	// sentence := TestAja()
	// fmt.Println(sentence)

	result := calculation.Add(8, 2)

	fmt.Println(result)
}
