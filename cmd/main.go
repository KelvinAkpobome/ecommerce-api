package main

import (
	"ecommerce-api/adapters/primary/http"
	"ecommerce-api/adapters/primary/http/middleware"
	"ecommerce-api/adapters/secondary/db"
	"ecommerce-api/application"
	"ecommerce-api/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dbConn := config.ConnectDB()

    sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatal("Failed to get *sql.DB from GORM:", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Failed to close database connection:", err)
		}
	}()

	// Repositories
	userRepo := db.NewUserRepository(dbConn)
	productRepo := db.NewProductRepository(dbConn)
	orderRepo := db.NewOrderRepository(dbConn)

	// Services
	userService := application.NewUserService(userRepo)
	productService := application.NewProductService(productRepo)
	orderService := application.NewOrderService(orderRepo)

	// Controllers
	authController := http.NewAuthController(userService)
	productController := http.NewProductController(productService)
	orderController := http.NewOrderController(orderService, productService, userService)

	r := gin.Default()

	// Public routes
	public := r.Group("/api")
	public.POST("/register", authController.Register)
	public.POST("/login", authController.Login)

	private := r.Group("/api")
	private.Use(middleware.AuthMiddleware())

	admin := private.Group("/admin")
	admin.Use(middleware.IsAdminMiddleware())
	admin.PUT("/products/:product_id", productController.UpdateProduct)
	admin.DELETE("/products/:product_id", productController.DeleteProduct)
	admin.POST("/products", productController.CreateProduct)
	admin.PUT("/orders/:order_id", orderController.UpdateOrderStatus)


	private.GET("/products", productController.GetAllProducts)
	private.GET("/products/:product_id", productController.GetProductByID)
	private.POST("/orders", orderController.PlaceOrder)
	private.GET("/orders/user", orderController.GetOrdersByUser)
	private.PUT("/orders/:order_id", orderController.CancelOrder)

	r.Run(":8080")
}
