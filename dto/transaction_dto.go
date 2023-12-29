package dto

type TransactionDTO struct {
	ID        uint64 `json:"id" form:"id"`
	ProductID uint64 `json:"productID" form:"productID" binding:"required"`
	StockID   uint64 `json:"stockID" form:"stockID" binding:"required"`
	Quantity  int    `json:"quantity,string" form:"quantity" binding:"required"`
}
