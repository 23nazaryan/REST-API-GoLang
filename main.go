package main

import (
	"gin/config"
	"gin/controllers"
	"gin/middleware"
	"gin/repositories"
	"gin/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                       = config.SetupDatabaseConnection()
	userRepository    repositories.UserRepository    = repositories.NewUserRepository(db)
	productRepository repositories.ProductRepository = repositories.NewProductRepository(db)
	stockRepository   repositories.StockRepository   = repositories.NewStockRepository(db)
	txnRepository     repositories.TxnRepository     = repositories.NewTxnRepository(db)
	jwtService        services.JWTService            = services.NewJWTService()
	userService       services.UserService           = services.NewUserService(userRepository)
	authService       services.AuthService           = services.NewAuthService(userRepository)
	productService    services.ProductService        = services.NewProductService(productRepository)
	stockService      services.StockService          = services.NewStockService(stockRepository)
	txnService        services.TxnService            = services.NewTxnService(txnRepository)
	authController    controllers.AuthController     = controllers.NewAuthController(authService, jwtService)
	userController    controllers.UserController     = controllers.NewUserController(userService, jwtService)
	productController controllers.ProductController  = controllers.NewProductController(productService, jwtService)
	stockController   controllers.StockController    = controllers.NewStockController(stockService, jwtService)
	txnController     controllers.TxnController      = controllers.NewTxnController(txnService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.GET("/verify", authController.Verify)
		authRoutes.POST("/check-hash", authController.CheckHash)
		authRoutes.POST("/activate", authController.Activate)
		authRoutes.POST("/forgot-password", authController.ForgotPassword)
		authRoutes.POST("/reset-password", authController.Activate)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.POST("/create", userController.Create)
		userRoutes.GET("/list", userController.FindAll)
		userRoutes.PUT("/update", userController.Update)
		userRoutes.DELETE("/delete/:id", userController.Delete)
	}

	productRoutes := r.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.POST("/create", productController.Create)
		productRoutes.PUT("/update", productController.Update)
		productRoutes.DELETE("/delete/:id", productController.Delete)
		productRoutes.GET("/list/:qtyFilter", productController.FindAll)
	}

	stockRoutes := r.Group("api/stock", middleware.AuthorizeJWT(jwtService))
	{
		stockRoutes.POST("/create", stockController.Create)
		stockRoutes.PUT("/update", stockController.Update)
		stockRoutes.DELETE("/delete/:id", stockController.Delete)
		stockRoutes.GET("/list/:type", stockController.FindAll)
	}

	transactionRoutes := r.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.POST("/create", txnController.Create)
		transactionRoutes.PUT("/update", txnController.Update)
		transactionRoutes.DELETE("/delete/:id", txnController.Delete)
		transactionRoutes.GET("/list/:section/:id", txnController.List)
	}

	err := r.Run()
	if err != nil {
		panic("Failed to create routes")
	}
}
