package storage

import (
	"errors"
	e "example/gotodo/entity"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MsgIncorrectSignUp = "sign up is failed"

	ErrConnectionFail   = errors.New("cannot connection to db")
	ErrCloseConnection  = errors.New("errors during attemption close db connection")
	ErrListNotFound     = errors.New("list not found")
	ErrWithGettingLists = errors.New("some trubles with getting lists")

	MsgListCannotUpdate = "list cannot updated"
	MsgListWasUpdated   = "list was updated"
	MsgListWasDeleted   = "list with id: %s was deleted"

	ErrTaskNotFound = errors.New("task not found")

	MsgTaskWasUpdated   = "task was updated"
	MsgTaskCannotUpdate = "task cannot updated"
	MsgTaskWasDeleted   = "task was deleted"

	LocalPgConnection  = "host=localhost user=userlst password=admin dbname=todo port=5432 sslmode=disable"
	DockerPgConnection = "host=my-postgres user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
)

type postgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage() (*postgresStorage, error) {
	db, err := gorm.Open(postgres.Open(DockerPgConnection), &gorm.Config{})
	if err != nil {
		return &postgresStorage{}, ErrConnectionFail
	}
	db.AutoMigrate(&e.TdList{})
	db.AutoMigrate(&e.TdTask{})
	db.AutoMigrate(&e.User{})
	return &postgresStorage{db: db}, nil
}

// USER

func (s *postgresStorage) CreateUser(user *e.User) (string, error) {
	result := s.db.Create(&user)
	if result.Error != nil {
		return MsgIncorrectSignUp, result.Error
	}
	return "success", nil
}

// LIST

func (s *postgresStorage) GetAll() ([]e.TdList, error) {
	var lists []e.TdList
	err := s.db.Model(&e.TdList{}).Find(&lists).Error
	if err != nil {
		return lists, ErrWithGettingLists
	}
	return lists, nil
}

func (s *postgresStorage) GetByID(id uuid.UUID) (e.TdList, error) {
	var list e.TdList
	err := s.db.Preload("Tasks").First(&list, id).Error
	if err != nil {
		return list, ErrListNotFound
	}
	return list, nil
}

func (s *postgresStorage) Create(list *e.TdList) (uuid.UUID, error) {
	result := s.db.Create(&list)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return list.Id, nil
}

func (s *postgresStorage) Update(list *e.TdList) (string, error) {
	var curList e.TdList
	err := s.db.First(&curList, &list.Id).Error
	if err != nil {
		return MsgListCannotUpdate, ErrListNotFound
	}
	curList.Name = list.Name
	curList.Date = list.Date
	s.db.Save(&curList)
	return MsgListWasUpdated, nil
}

func (s *postgresStorage) Delete(id uuid.UUID) (string, error) {
	s.db.Where("list_id = ?", id).Delete(&e.TdTask{})
	s.db.Where("id = ?", id).Delete(&e.TdList{})
	result := fmt.Sprintf(MsgListWasDeleted, id.String())
	return result, nil
}

// TASK

func (s *postgresStorage) GetTaskByID(id uuid.UUID) (e.TdTask, error) {
	var task e.TdTask
	err := s.db.First(&task, id).Error
	if err != nil {
		return task, ErrTaskNotFound
	}
	return task, nil
}

func (s *postgresStorage) CreateTask(task *e.TdTask) (uuid.UUID, error) {
	result := s.db.Create(&task)
	if result.Error != nil {
		return uuid.Nil, result.Error
	}
	return task.Id, nil
}

func (s *postgresStorage) UpdateTask(task *e.TdTask) (string, error) {
	var curTask e.TdTask
	err := s.db.First(&curTask, &task.Id).Error
	if err != nil {
		return MsgTaskCannotUpdate, ErrTaskNotFound
	}
	curTask.Name = task.Name
	curTask.Description = task.Description
	curTask.Date = task.Date
	s.db.Save(&curTask)
	return MsgTaskWasUpdated, nil
}

func (s *postgresStorage) DeleteTask(id uuid.UUID) (string, error) {
	s.db.Where("id = ?", id).Delete(&e.TdTask{})
	result := fmt.Sprintf(MsgTaskWasDeleted, id.String())
	return result, nil
}
