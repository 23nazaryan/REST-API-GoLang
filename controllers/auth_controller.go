package controllers

import (
	"fmt"
	"gin/dto"
	"gin/entities"
	"gin/helpers"
	"gin/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Verify(ctx *gin.Context)
	CheckHash(ctx *gin.Context)
	Activate(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entities.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generatedToken
		response := helpers.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helpers.BuildErrorResponse("Please check again your credentials", "Invalid credentials", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Verify(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.authService.FindByID(id)
	user.Token = token.Raw
	res := helpers.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) CheckHash(ctx *gin.Context) {
	var hashDTO dto.HashDTO
	errDTO := ctx.ShouldBind(&hashDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user := c.authService.VerifyHash(hashDTO)
	if user != nil {
		response := helpers.BuildResponse(true, "OK!", helpers.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helpers.BuildErrorResponse("Please check again your credentials", "The link is out of date", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Activate(ctx *gin.Context) {
	var pwdDTO dto.PwdDTO
	errDTO := ctx.ShouldBind(&pwdDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user := c.authService.SetPassword(pwdDTO)
	generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))
	user.Token = generatedToken
	res := helpers.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) ForgotPassword(ctx *gin.Context) {
	var forgotDTO dto.ForgotDTO
	errDTO := ctx.ShouldBind(&forgotDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err := c.authService.SendForgotEmail(forgotDTO.Email)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to sending email process", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "OK!", helpers.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
