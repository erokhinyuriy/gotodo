package userservice

import (
	e "example/gotodo/entity"
)

type UserRepository interface {
	CreateUser(user *e.User) (string, error)
	GetUser(email string) (e.User, error)
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

func (u *service) GetUser(email string) (e.User, error) {
	return u.repo.GetUser(email)
}
