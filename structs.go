package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthDate string
	createdOn time.Time

} 

// Below is a Constructor function
func newUser(firstName, lastName, birthDate string) (*user, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birth date are all required")
	} 


	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdOn: time.Now(),

}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	//var appUser *user
	
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	

	appUser.displayUserDetails()
	appUser.clearUserData()
	appUser.displayUserDetails()


}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// ... do something awesome with that gathered data!

// Below this are methods
func (details *user) displayUserDetails() {
	fmt.Println(details.firstName, details.lastName, details.birthDate)

}

func (details *user) clearUserData() {
	details.firstName = ""
	details.lastName = ""
}
