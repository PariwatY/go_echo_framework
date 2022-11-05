package services

import (
	"errors"
	model "go_ktb_test/model/product"
	"go_ktb_test/repositories"

	"github.com/go-redis/redis/v8"
)

type catalogService struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogService(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogService{productRepo, redisClient}
}

func (s catalogService) GetProducts() (products []Product, err error) {

	//Get Product Data
	productDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	//Append productDB to products
	for _, p := range productDB {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	return products, err
}

func (s catalogService) GetProductById(id int) (product Product, err error) {

	productDB, err := s.productRepo.GetProductById(id)
	if err != nil {
		return product, err
	}

	product = Product{
		ID:       productDB.ID,
		Name:     productDB.Name,
		Quantity: productDB.Quantity,
	}

	return product, err
}

func (s catalogService) SaveProduct(name string, quantity int) (product Product, err error) {
	productDB, err := s.productRepo.SaveProduct(name, quantity)

	if err != nil {
		return Product{}, err
	}

	product = Product{
		ID:       productDB.ID,
		Name:     productDB.Name,
		Quantity: productDB.Quantity,
	}

	return product, err
}

func (s catalogService) UpdateProductById(productRequest model.ProductRequest) (product Product, err error) {

	productDB, err := s.productRepo.UpdateProductById(productRequest)

	if productDB.Name == "" || err != nil {
		return product, errors.New("Product not found")
	}

	product = Product{
		ID:       productDB.ID,
		Name:     productDB.Name,
		Quantity: productDB.Quantity,
	}

	return product, err
}

func (s catalogService) DeleteProductById(id int) (err error) {
	err = s.productRepo.DeleteProductById(id)
	return err
}
