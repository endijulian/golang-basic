package main

import "fmt"

func main() {

	sentence := printMyResult("Saya Sedang ")
	fmt.Println(sentence)

	valueReturn := add(10, 20)
	fmt.Println(valueReturn)
}

// func printMyResult() {
// 	fmt.Println("Saya sedang belajar golang")
// }

// func printMyResult(sentence string) {
// 	fmt.Println(sentence)
// }

func printMyResult(sentence string) string {
	newSentence := sentence + "Belajar Golang"

	return newSentence
}

// func add(number int, numberTwo int) int {
// 	// result := number + numberTwo
// 	// return result

// 	return number + numberTwo
// }

func add(number, numberTwo int) int {
	// result := number + numberTwo
	// return result

	return number + numberTwo
}
