package main

import (
	"go_ktb_test/config"
	"go_ktb_test/handlers"
	"go_ktb_test/repositories"
	"go_ktb_test/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initial Config Database Connection
	db := config.InitialDB()

	// Initial Config Redis Connection
	redisClient := config.InitRedis()
	_ = redisClient

	// Initial Product Repository
	productRepo := repositories.NewProductRepositoryDB(db, redisClient)
	// Initial Product Service
	productService := services.NewCatalogService(productRepo, redisClient)
	// productService := services.NewCatalogServiceRedis(productRepo, redisClient)

	// Initial Product Handler
	productHandler := handlers.NewCatalogHandler(productService)

	app := echo.New()

	// Create, Insert product
	app.POST("product", productHandler.SaveProduct)
	// Get All product
	app.GET("products", productHandler.GetProducts)
	// Get product by Id
	app.GET("product/:id", productHandler.GetProductById)
	// Update product
	app.PUT("product", productHandler.UpdateProductById)
	// Delete product
	app.DELETE("product/:id", productHandler.DeleteProductById)
	// Call Another
	app.GET("pokemons", handlers.GetPokeMonList)

	if err := app.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
