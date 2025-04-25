package taskservice

import (
	e "example/gotodo/entity"

	"github.com/google/uuid"
)

type TaskRepository interface {
	GetTaskByID(id uuid.UUID) (e.TdTask, error)
	CreateTask(task *e.TdTask) (uuid.UUID, error)
	UpdateTask(task *e.TdTask) (string, error)
	DeleteTask(id uuid.UUID) (string, error)
}

type service struct {
	repo TaskRepository
}

func New(repo TaskRepository) *service {
	return &service{repo: repo}
}

func (t *service) GetTaskByID(id uuid.UUID) (e.TdTask, error) {
	return t.repo.GetTaskByID(id)
}

func (t *service) CreateTask(task *e.TdTask) (uuid.UUID, error) {
	return t.repo.CreateTask(task)
}

func (t *service) UpdateTask(task *e.TdTask) (string, error) {
	return t.repo.UpdateTask(task)
}

func (t *service) DeleteTask(id uuid.UUID) (string, error) {
	return t.repo.DeleteTask(id)
}
