package entities

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key:auto_increment" json:"ID"`
	SkuID    string `gorm:"type:varchar(50);not null" json:"skuID"`
	Title    string `gorm:"type:varchar(50);not null" json:"title"`
	Box      string `gorm:"type:varchar(50)" json:"box"`
	Min      int    `gorm:"type:integer(11);not null" json:"min"`
	Max      int    `gorm:"type:integer(11);not null" json:"max"`
	Quantity int    `gorm:"type:integer(11);default:0;not null" json:"quantity"`
}
