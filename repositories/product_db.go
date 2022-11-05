package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"go_ktb_test/constant"
	model "go_ktb_test/model/product"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryDB(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	//Create Auto Table products
	db.AutoMigrate(&product{})

	return productRepositoryDB{db, redisClient}
}

func (r productRepositoryDB) GetProducts() (products []product, err error) {

	key := "repository::GetProducts"
	//Get Product from redis
	productJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productJson), &products)
		if err == nil {
			fmt.Println("data from redis")
			return products, nil
		}
	}

	//If Get product from redis not found then Get from Database to set data to redis
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	//Marshal product data to byte for preparing data to redis
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	//Set product data from database to redis for 10 second
	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("data from database")
	return products, err

}

func (r productRepositoryDB) GetProductById(id int) (productById product, err error) {
	//Search product data by id
	err = r.db.Where("id = ?", id).Find(&productById).Error

	//Check product data if not exist will return error
	if len(productById.Name) == 0 {
		return productById, constant.ErrNotFound
	}
	return productById, err
}

func (r productRepositoryDB) SaveProduct(name string, quantity int) (productData product, err error) {
	//Search product for verify duplicate data by name
	r.db.Where("name = ?", name).Find(&productData)

	if productData.Name == name {
		return product{}, constant.ErrDuplicate
	}

	productData = product{
		Name:     name,
		Quantity: quantity,
	}

	//Created product data
	tx2 := r.db.Create(&productData)
	if tx2.Error != nil {
		return
	}

	fmt.Println("Product created ", productData)
	return productData, err
}

func (r productRepositoryDB) UpdateProductById(productRequest model.ProductRequest) (productById product, err error) {
	//Search product for update data by id
	tx := r.db.Where("id = ?", productRequest.ID).Find(&productById)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Before Update: ", productById)

	productById.Name = productRequest.Name
	productById.Quantity = productRequest.Quantity

	//Update product data
	tx2 := r.db.Save(&productById)
	if tx2.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println("After Update: ", productById)
	return productById, err
}

func (r productRepositoryDB) DeleteProductById(id int) (err error) {

	productId := &product{}

	//Search product for delete data by id
	tx := r.db.Where("id = ?", id).Find(productId)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	//Delete product for delete data by id
	tx2 := r.db.Where("id = ?", id).Delete(productId)
	if tx2.Error != nil {
		fmt.Println(tx2.Error)
		return
	}

	return err
}
