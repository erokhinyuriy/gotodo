package entity

import (
	"time"

	"github.com/google/uuid"
)

type TdList struct {
	Id     uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserId uuid.UUID `json:"user_id"`
	Name   string
	Date   time.Time
	Tasks  []TdTask `gorm:"foreignKey:ListId"`
}

type Tabler interface {
	TableName() string
}

func (TdList) TableName() string {
	return "td_list"
}
