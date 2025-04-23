package entity

import (
	"time"

	"github.com/google/uuid"
)

type TdList struct {
	Id   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name string
	Date time.Time
}

type Tabler interface {
	TableName() string
}

func (TdList) TableName() string {
	return "td_list"
}
