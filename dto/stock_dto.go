package dto

type StockDTO struct {
	ID      uint64 `json:"id" form:"id"`
	Title   string `json:"title" form:"title" binding:"required"`
	Address string `json:"address" form:"address"`
	Type    string `json:"type" form:"type" binding:"required"`
}
