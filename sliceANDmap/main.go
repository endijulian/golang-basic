package main

import "fmt"

func main() {
	student := []map[string]string{
		{"Name": "Endi", "Score": "A"},
		{"Name": "Julian", "Score": "B"},
		{"Name": "Array", "Score": "C"},
	}

	// fmt.Println(student)

	for _, value := range student {
		fmt.Println(value["Name"], "--", value["Score"])
	}
}
