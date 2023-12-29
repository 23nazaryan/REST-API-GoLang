package dto

type ProductDTO struct {
	ID       uint64 `json:"id" form:"id"`
	SkuID    string `json:"skuID" form:"sku_id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Box      string `json:"box" form:"box"`
	Min      int    `json:"min,string" form:"min" binding:"required"`
	Max      int    `json:"max,string" form:"max" binding:"required"`
	Quantity int    `json:"quantity" form:"quantity"`
}
