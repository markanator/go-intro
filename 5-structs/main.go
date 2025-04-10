package main

import (
	"example.com/markanator/structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please enter your first name")
	userLastName := getUserData("Please enter your last name")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY)")

	var appUser *user.User
	appUser, err := user.New(
		userFirstName,
		userLastName,
		userBirthdate,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use the user struct method
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(question string) string {
	fmt.Printf(question)
	var field string
	_, err := fmt.Scanln(&field)
	if err != nil {
		return ""
	}
	return field
}
