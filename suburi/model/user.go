package model

import "errors"
import "github.com/google/uuid"

type User struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	Description string `json:"description"`
	PhotoURL    string `json:"photoURL"`
	Email       string `json:"email"`
}

func NewUser(userName, description, photoURL, email string) (*User, error) {
	if userName == "" {
		return nil, errors.New("userName is empty")
	}
	if description == "" {
		return nil, errors.New("description is empty")
	}
	if photoURL == "" {
		return nil, errors.New("photoURL is empty")
	}
	if email == "" {
		return nil, errors.New("email is empty")
	}

	return &User{
		UserId:      uuid.New().String(),
		UserName:    userName,
		Description: description,
		PhotoURL:    photoURL,
		Email:       email}, nil
}

func (user *User) Equals(other User) bool {
	return user.UserId == other.UserId
}
