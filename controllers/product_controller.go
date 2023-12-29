package controllers

import (
	"gin/dto"
	"gin/helpers"
	"gin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type productController struct {
	productService services.ProductService
	jwtService     services.JWTService
}

func NewProductController(productService services.ProductService, jwtService services.JWTService) ProductController {
	return &productController{
		productService: productService,
		jwtService:     jwtService,
	}
}

func (c *productController) Create(ctx *gin.Context) {
	var productDTO dto.ProductDTO
	errDTO := ctx.ShouldBind(&productDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.productService.IsDuplicateSkuID(productDTO.SkuID) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate SKU ID", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdProduct := c.productService.Create(productDTO)
		response := helpers.BuildResponse(true, "OK!", createdProduct)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *productController) Update(ctx *gin.Context) {
	var productDTO dto.ProductDTO
	errDTO := ctx.ShouldBind(&productDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	product := c.productService.Update(productDTO)
	res := helpers.BuildResponse(true, "OK!", product)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	result := c.productService.Delete(id)
	if result != nil {
		res := helpers.BuildErrorResponse("Failed to process request", result.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "OK!", helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) FindAll(ctx *gin.Context) {
	qtyFilter := ctx.Param("qtyFilter")
	products := c.productService.FindAll(qtyFilter)
	res := helpers.BuildResponse(true, "OK!", products)
	ctx.JSON(http.StatusOK, res)
}
