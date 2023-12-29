package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID        uint64 `gorm:"primary_key:auto_increment" json:"ID"`
	ProductID uint64 `gorm:"type:integer(11);not null" json:"productID"`
	StockID   uint64 `gorm:"type:integer(11);not null" json:"stockID"`
	Quantity  int    `gorm:"type:integer(11);default:0;not null" json:"quantity"`
	Product   Product
	Stock     Stock
}
