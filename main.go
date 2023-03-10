package main

import (
	"e-commerce-api/config"
	_orderData "e-commerce-api/feature/order/data"
	_orderHandler "e-commerce-api/feature/order/handler"
	_orderService "e-commerce-api/feature/order/service"
	_productData "e-commerce-api/feature/product/data"
	_productHandler "e-commerce-api/feature/product/handler"
	_productService "e-commerce-api/feature/product/service"
	"e-commerce-api/feature/users/data"
	"e-commerce-api/feature/users/handler"
	"e-commerce-api/feature/users/services"

	cData "e-commerce-api/feature/cart/data"
	cHandler "e-commerce-api/feature/cart/handler"
	cService "e-commerce-api/feature/cart/service"

	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	// panggil fungsi Migrate untuk buat table baru di database
	config.Migrate(db)

	v := validator.New()
	cld := config.NewCloudinary(*cfg)

	productData := _productData.New(db)
	productService := _productService.New(productData, v, cld)
	productHandler := _productHandler.New(productService)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	orderData := _orderData.New(db)
	orderService := _orderService.New(orderData)
	orderHandler := _orderHandler.New(orderService)
	cartData := cData.New(db)
	cartSrv := cService.New(cartData)
	cartHdl := cHandler.New(cartSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/products", productHandler.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products", productHandler.GetAll())
	e.GET("/products/:product_id", productHandler.GetByID())
	e.PUT("/products/:product_id", productHandler.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/products/:product_id", productHandler.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/orders", orderHandler.Create(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/orders", orderHandler.GetAll(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/orders/:order_id/cancel", orderHandler.Cancel(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/orders/:order_id/confirm", orderHandler.Confirm(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/orders/buy/:order_id", orderHandler.GetOrderBuy(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/orders/sell/:order_id", orderHandler.GetOrderSell(), middleware.JWT([]byte(config.JWT_KEY)))
	e.POST("/orders/callback", orderHandler.Callback())

	e.POST("/carts/:product_id", cartHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/carts", cartHdl.GetAll(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/carts/:cart_id", cartHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/carts/:cart_id", cartHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}

}
