package http

import (
	"net/http"
	"strconv"

	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/model"
	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(uc usecase.UseCase) *gin.Engine {
	r := gin.Default()

	r.POST("/products", func(c *gin.Context) {
		var p model.Product
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := uc.AddProduct(&p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		p, err := uc.GetProduct(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, p)
	})

	r.PATCH("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var p model.Product
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p.ID = id
		if err := uc.UpdateProduct(&p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := uc.DeleteProduct(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})

	r.GET("/products", func(c *gin.Context) {
		products, err := uc.ListProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	return r
}
