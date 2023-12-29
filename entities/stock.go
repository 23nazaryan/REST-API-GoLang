package entities

import (
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	ID      uint64 `gorm:"primary_key:auto_increment" json:"ID"`
	Title   string `gorm:"type:varchar(50);not null" json:"title"`
	Address string `gorm:"type:varchar(255)" json:"address"`
	Type    string `gorm:"type:varchar(10);not null" json:"type"`
}
