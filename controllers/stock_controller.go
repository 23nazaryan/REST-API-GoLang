package controllers

import (
	"gin/dto"
	"gin/helpers"
	"gin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StockController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type stockController struct {
	stockService services.StockService
	jwtService   services.JWTService
}

func NewStockController(stockService services.StockService, jwtService services.JWTService) StockController {
	return &stockController{
		stockService: stockService,
		jwtService:   jwtService,
	}
}

func (c *stockController) Create(ctx *gin.Context) {
	var stockDTO dto.StockDTO
	errDTO := ctx.ShouldBind(&stockDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	createdStock := c.stockService.Create(stockDTO)
	response := helpers.BuildResponse(true, "OK!", createdStock)
	ctx.JSON(http.StatusOK, response)
}

func (c *stockController) Update(ctx *gin.Context) {
	var stockDTO dto.StockDTO
	errDTO := ctx.ShouldBind(&stockDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	stock := c.stockService.Update(stockDTO)
	res := helpers.BuildResponse(true, "OK!", stock)
	ctx.JSON(http.StatusOK, res)
}

func (c *stockController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	result := c.stockService.Delete(id)
	if result != nil {
		res := helpers.BuildErrorResponse("Failed to process request", result.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "OK!", helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *stockController) FindAll(ctx *gin.Context) {
	stockType := ctx.Param("type")
	stocks := c.stockService.FindAll(stockType)
	res := helpers.BuildResponse(true, "OK!", stocks)
	ctx.JSON(http.StatusOK, res)
}
