package models

import (
	"errors"
	"fmt"
)

// User data models f
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Users var initialization
var (
	Users  []*User
	nextID = 1
)

// GetUsers returns all the users in variable Users
func GetUsers() []*User {
	return Users
}

// AddUser append a user in variable Users
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	Users = append(Users, &u)
	return u, nil
}

// GetUserByID allows to get user based on provided id
func GetUserByID(id int) (User, error) {
	for _, u := range Users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser allows to update user based on provided id
func UpdateUser(u User) (User, error) {
	for i, candidate := range Users {
		if candidate.ID == u.ID {
			Users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID allows to remove user based on provided id
func RemoveUserByID(id int) error {
	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
