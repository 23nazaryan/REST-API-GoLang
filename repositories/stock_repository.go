package repositories

import (
	"gin/entities"
	"gorm.io/gorm"
)

type StockRepository interface {
	Insert(stock entities.Stock) entities.Stock
	Update(stock entities.Stock) entities.Stock
	Delete(stockID string) error
	FindAll(stockType string) []entities.Stock
}

type stockConnection struct {
	connection *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockConnection{
		connection: db,
	}
}

func (db *stockConnection) Insert(stock entities.Stock) entities.Stock {
	db.connection.Save(&stock)
	return stock
}

func (db *stockConnection) Update(stock entities.Stock) entities.Stock {
	db.connection.Save(&stock)
	return stock
}

func (db *stockConnection) Delete(stockID string) error {
	var stock entities.Stock
	db.connection.Find(&stock, stockID)
	res := db.connection.Delete(&stock)
	return res.Error
}

func (db *stockConnection) FindAll(stockType string) []entities.Stock {
	var stocks []entities.Stock

	switch stockType {
	case "both":
		db.connection.Find(&stocks)
	case "fiscal":
		db.connection.Where("type = ?", stockType).Find(&stocks)
	case "non-fiscal":
		db.connection.Where("type = ?", stockType).Find(&stocks)
	}

	return stocks
}
