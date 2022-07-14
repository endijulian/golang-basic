package main

import "fmt"

func main() {
	// var myMap map[string]int

	// myMap = map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["Go"] = 9
	// myMap["PHP"] = 8
	// myMap["Go"] = 10

	// fmt.Println(myMap)

	myMap := map[string]string{
		"Ruby":      "Is beautiful",
		"Go":        "Is super fast",
		"Javasript": "hype",
	}

	for key, value := range myMap {
		fmt.Println("Key :", key, "Value :", value)
	}

	fmt.Println("========================")

	delete(myMap, "Ruby")

	for key, value := range myMap {
		fmt.Println("Key :", key, "Value :", value)
	}

}
