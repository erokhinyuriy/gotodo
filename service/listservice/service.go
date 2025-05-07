package listservice

import (
	e "example/gotodo/entity"

	"github.com/google/uuid"
)

type ListRepository interface {
	GetByID(id uuid.UUID, uid uuid.UUID) (e.TdList, error)
	GetAll(uid uuid.UUID) ([]e.TdList, error)
	Create(list *e.TdList) (uuid.UUID, error)
	Update(list *e.TdList) (string, error)
	Delete(id uuid.UUID) (string, error)
}

type service struct {
	repo ListRepository
}

func New(repo ListRepository) *service {
	return &service{repo: repo}
}

func (lst *service) GetAll(uid uuid.UUID) ([]e.TdList, error) {
	return lst.repo.GetAll(uid)
}

func (lst *service) GetByID(id uuid.UUID, uid uuid.UUID) (e.TdList, error) {
	return lst.repo.GetByID(id, uid)
}

func (lst *service) Create(list *e.TdList) (uuid.UUID, error) {
	return lst.repo.Create(list)
}

func (lst *service) Update(list *e.TdList) (string, error) {
	return lst.repo.Update(list)
}

func (lst *service) Delete(id uuid.UUID) (string, error) {
	return lst.repo.Delete(id)
}
