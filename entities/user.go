package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key:auto_increment" json:"ID"`
	Name     string `gorm:"type:varchar(50);not null" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(50);not null" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Role     string `gorm:"type:varchar(50);not null" json:"role"`
	Hash     string `gorm:"type:varchar(255)" json:"hash"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
