package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Endi"
	user.LastName = "Julian"
	user.Email = "endi@gmail.com"
	user.IsActive = true

	fmt.Println(user)
}

func displayUser(user string) string {
	return "Name : Endi Julian, Email : endi@gmail.com"
}
