package services

import (
	model "go_ktb_test/model/product"
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type CatalogService interface {
	// get all product
	GetProducts() ([]Product, error)

	// get product by id
	GetProductById(id int) (Product, error)

	//save product
	SaveProduct(name string, quantity int) (Product, error)

	//update product
	UpdateProductById(productRequest model.ProductRequest) (Product, error)

	//delete product by id
	DeleteProductById(id int) error
}
