package dao

import (
	"errors"
	"github.com/laqiiz/graphql-go-learning/suburi/model"
)

type UserRepository struct {
	users []*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// store event to dao
func (r *UserRepository) Store(user *model.User) {
	r.users = append(r.users, user)
}

func (r UserRepository) FindById(userId string) (*model.User, error) {
	for _, u := range r.users {
		if u.UserId == userId {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r UserRepository) List() []*model.User {
	return r.users
}
