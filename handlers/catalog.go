package handlers

import (
	"github.com/labstack/echo/v4"
)

type CatalogHandler interface {
	//get all product
	GetProducts(c echo.Context) error

	//get product by id
	GetProductById(c echo.Context) error

	//save product
	SaveProduct(c echo.Context) error

	//update product
	UpdateProductById(c echo.Context) error

	//delete product by id
	DeleteProductById(c echo.Context) error
}
