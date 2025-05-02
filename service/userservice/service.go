package userservice

import (
	e "example/gotodo/entity"
)

type UserRepository interface {
	CreateUser(user *e.User) (string, error)
}

type service struct {
	repo UserRepository
}

func New(repo UserRepository) *service {
	return &service{repo: repo}
}

func (u *service) CreateUser(user *e.User) (string, error) {
	return u.repo.CreateUser(user)
}
