package entity

import (
	"time"

	"github.com/google/uuid"
)

type TdTask struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ListId      uuid.UUID `json:"list_id"`
	Name        string
	Description string
	Date        time.Time
}

func (TdTask) TableName() string {
	return "td_task"
}
