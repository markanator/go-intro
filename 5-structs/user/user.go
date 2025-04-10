package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

// inherits from User by anonymous embedding
// we can call User methods off the rip without chaining
type Admin struct {
	email    string
	password string
	User
}

// adding methods to a struct
// access to a user struct using a receiver
func (u *User) OutputUserDetails() {
	fmt.Printf("%s, %s born on %s\n", u.lastName, u.firstName, u.birthday)
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Mark",
			lastName:  "Admin",
			birthday:  "---",
			createdAt: time.Now(),
		},
	}
}

func New(first, last, birthday string) (*User, error) {
	if first == "" || last == "" || birthday == "" {
		return nil, errors.New("invalid inputs")
	}
	return &User{
		firstName: first,
		lastName:  last,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}
