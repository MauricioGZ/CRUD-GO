package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	user := e.Group("/user")
	address := user.Group("/address")
	categories := e.Group("/categories")
	products := e.Group("/products")
	order := user.Group("/order")
	orderItem := order.Group("/item")

	user.POST("/register", a.RegisterUser)
	user.GET("/login", a.LoginUser)

	address.GET("", a.GetAddresses)
	address.POST("/register", a.AddAddress)
	address.PATCH("/update", a.UpdateAddress)
	address.DELETE("/delete/:id", a.DeleteAddress)

	categories.POST("/register", a.RegisterCategory)
	categories.GET("/", a.GetAllCategories)

	products.POST("/add", a.AddProduct)
	products.GET("", a.GetAllProducts)
	products.GET("/:id", a.GetProductByID)
	products.PATCH("/:id", a.UpdateProductByID)
	products.DELETE("/:id", a.DeleteProductByID)

	order.POST("", a.AddOrder)
	order.GET("", a.GetOrders)

	orderItem.GET("", a.GetOrderItemsByUser)
}
