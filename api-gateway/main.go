package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Проверка работоспособности
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Маршруты для Inventory Service
	r.POST("/products", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8081/products")
	})
	r.GET("/products/:id", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8081/products/"+c.Param("id"))
	})
	r.PATCH("/products/:id", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8081/products/"+c.Param("id"))
	})
	r.DELETE("/products/:id", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8081/products/"+c.Param("id"))
	})
	r.GET("/products", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8081/products")
	})

	// Маршруты для Order Service
	r.POST("/orders", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8082/orders")
	})
	r.GET("/orders/:id", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8082/orders/"+c.Param("id"))
	})
	r.PATCH("/orders/:id", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8082/orders/"+c.Param("id"))
	})
	r.GET("/orders", func(c *gin.Context) {
		c.Redirect(307, "http://localhost:8082/orders")
	})

	r.Run(":8080") // Запуск на порту 8080
}
