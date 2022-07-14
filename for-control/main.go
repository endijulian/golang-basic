package main

import "fmt"

func main() {

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("Endi Julian : ", i)
	// }

	// i := 1
	// for i <= 50 {
	// 	fmt.Println("Looping sampai 50 : ", i)
	// 	i++
	// }

	// title := "Golang the best language"

	// for index, letter := range title {
	// 	fmt.Println("index :", index, "letter :", string(letter))
	// }

	title := "Golang the best language"

	for _, letter := range title {
		fmt.Println("letter :", string(letter))
	}
}
