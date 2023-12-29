package controllers

import (
	"gin/dto"
	"gin/helpers"
	"gin/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TxnController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

type txnController struct {
	txnService services.TxnService
	jwtService services.JWTService
}

func NewTxnController(txnService services.TxnService, jwtService services.JWTService) TxnController {
	return &txnController{
		txnService: txnService,
		jwtService: jwtService,
	}
}

func (c *txnController) Create(ctx *gin.Context) {
	var txnDTO dto.TransactionDTO
	errDTO := ctx.ShouldBind(&txnDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	createdTxn := c.txnService.Create(txnDTO)
	response := helpers.BuildResponse(true, "OK!", createdTxn)
	ctx.JSON(http.StatusOK, response)
}

func (c *txnController) Update(ctx *gin.Context) {
	var txnDTO dto.TransactionDTO
	errDTO := ctx.ShouldBind(&txnDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	txn := c.txnService.Update(txnDTO)
	res := helpers.BuildResponse(true, "OK!", txn)
	ctx.JSON(http.StatusOK, res)
}

func (c *txnController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	result := c.txnService.Delete(id)
	if result != nil {
		res := helpers.BuildErrorResponse("Failed to process request", result.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.BuildResponse(true, "OK!", helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (c *txnController) List(ctx *gin.Context) {
	section := ctx.Param("section")
	id := ctx.Param("id")
	transactions := c.txnService.List(section, id)
	res := helpers.BuildResponse(true, "OK!", transactions)
	ctx.JSON(http.StatusOK, res)
}
