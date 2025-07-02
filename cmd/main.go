package main

import (
	"product_stock/controller"
	"product_stock/db"
	"product_stock/repository"
	"product_stock/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.POST("/product", productController.CreateProduct)
	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productId", productController.GetProductById)
	server.PUT("/product/:productId", productController.UpdateProduct)
	server.DELETE("/product/:productId", productController.DeleteProduct)

	server.Run(":8000")
}
