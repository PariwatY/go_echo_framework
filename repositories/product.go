package repositories

import (
	model "go_ktb_test/model/product"
)

type product struct {
	ID       int
	Name     string
	Quantity int
}

type ProductRepository interface {
	// get all product
	GetProducts() ([]product, error)

	// get product by id
	GetProductById(id int) (product, error)

	//save product
	SaveProduct(name string, quantity int) (product, error)

	//update product
	UpdateProductById(productRequest model.ProductRequest) (product, error)

	//delete product by id
	DeleteProductById(id int) error
}

// func mockData(db *gorm.DB) error {

// 	var count int64
// 	db.Model(&product{}).Count(&count)

// 	if count > 0 {
// 		return nil
// 	}

// 	products := []product{}
// 	seed := rand.NewSource(time.Now().UnixNano())
// 	random := rand.New(seed)
// 	for i := 0; i < 10; i++ {
// 		products = append(products, product{
// 			Name:     fmt.Sprintf("Product%v", i+1),
// 			Quantity: random.Int(),
// 		})
// 	}

// 	return db.Create(&products).Error
// }
