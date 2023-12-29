package repositories

import (
	"gin/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(product entities.Product) entities.Product
	Update(product entities.Product) entities.Product
	Delete(productID string) error
	FindAll(qtyFilter string) []entities.Product
	IsDuplicateSkuID(skuID string) (tx *gorm.DB)
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) Insert(product entities.Product) entities.Product {
	db.connection.Save(&product)
	return product
}

func (db *productConnection) Update(product entities.Product) entities.Product {
	db.connection.Save(&product)
	return product
}

func (db *productConnection) Delete(productID string) error {
	var product entities.Product
	db.connection.Find(&product, productID)
	res := db.connection.Delete(&product)
	return res.Error
}

func (db *productConnection) FindAll(qtyFilter string) []entities.Product {
	var products []entities.Product

	switch qtyFilter {
	case "both":
		db.connection.Find(&products)
	case "bellow":
		db.connection.Where("min > quantity").Find(&products)
	case "above":
		db.connection.Where("min < quantity").Find(&products)
	}

	return products
}

func (db *productConnection) IsDuplicateSkuID(skuID string) (tx *gorm.DB) {
	var product entities.Product
	return db.connection.Where("sku_id = ?", skuID).Take(&product)
}
