package main

import "fmt"

func main() {
	// name := "Golang"

	// var languages [5]string

	// languages[0] = "Go"
	// languages[1] = "Vue"
	// languages[2] = "Nuxt"
	// languages[3] = "Tailwind"
	// languages[4] = "Laravel"

	// languages := [5]string{"Go", "Nuxt", "Vue", "Laravel", "Tailwind"}
	// languages := [5]string{
	// 	"Go",
	// 	"Nuxt",
	// 	"Vue",
	// 	"Laravel",
	// 	"Tailwind",
	// }

	languages := [...]string{
		"Go",
		"Nuxt",
		"Vue",
		"Laravel",
		"Tailwind",
		"Ruby",
	}

	for index, lang := range languages {
		fmt.Println("Index :", index, "Language :", lang)
	}

	// fmt.Println(languages)

	// lenght := len(languages)

	// fmt.Println(lenght)
}
