package repository

import (
	"errors"
	"github.com/laqiiz/graphql-go-learning/suburi/model"
)

type UserRepository struct {
	users []*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{[]*model.User{}}
}

// store event to repository
func (r *UserRepository) Store(user *model.User) *UserRepository {
	r.users = append(r.users, user)
	return r
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
