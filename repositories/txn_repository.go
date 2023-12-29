package repositories

import (
	"gin/entities"
	"gorm.io/gorm"
	"log"
)

type TxnRepository interface {
	Insert(txn entities.Transaction) entities.Transaction
	Update(txn entities.Transaction) entities.Transaction
	Delete(txnID string) error
	List(section string, id string) []entities.Transaction
}

type txnConnection struct {
	connection *gorm.DB
}

func NewTxnRepository(db *gorm.DB) TxnRepository {
	return &txnConnection{
		connection: db,
	}
}

func (db *txnConnection) Insert(txn entities.Transaction) entities.Transaction {
	var product entities.Product
	db.connection.Find(&product, txn.ProductID)
	product.Quantity += txn.Quantity
	db.connection.Save(&product)
	var existsTxn entities.Transaction
	res := db.connection.Where("product_id = ? AND stock_id = ?", txn.ProductID, txn.StockID).Take(&existsTxn)

	if res.Error == nil {
		log.Println("res")
		existsTxn.Quantity += txn.Quantity
		db.connection.Save(&existsTxn)
		return existsTxn
	}

	db.connection.Save(&txn)
	return txn
}

func (db *txnConnection) Update(txn entities.Transaction) entities.Transaction {
	var oldTxn entities.Transaction
	db.connection.Find(&oldTxn, txn.ID)
	var product entities.Product
	db.connection.Find(&product, txn.ProductID)
	product.Quantity -= oldTxn.Quantity
	product.Quantity += txn.Quantity
	db.connection.Save(&product)
	db.connection.Save(&txn)
	return txn
}

func (db *txnConnection) Delete(txnID string) error {
	var transaction entities.Transaction
	db.connection.Find(&transaction, txnID)
	res := db.connection.Delete(&transaction)
	return res.Error
}

func (db *txnConnection) List(section string, id string) []entities.Transaction {
	var transactions []entities.Transaction

	switch section {
	case "product":
		db.connection.Where("product_id = ?", id).Preload("Product").Preload("Stock").Find(&transactions)
	case "stock":
		db.connection.Where("stock_id = ?", id).Preload("Product").Preload("Stock").Find(&transactions)
	}

	return transactions
}
