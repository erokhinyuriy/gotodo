package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string
	Email    string `gorm:"unique"`
	Password string
}
