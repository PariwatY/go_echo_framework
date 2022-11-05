# About Project
- Implement API with golang with Echo

# Requirements
- [x] go version 1.18
- [x] REST API (create, update, insert, delete data)
- [x] Storage: Redis and MySQL database
- [x] Struct (Object Model)
- [X] Call another REST API
- [x] Explain functions in source code
- [x] Postman API document


# Step 1 Preparing Docker image
```
docker pull redis
```
```
docker pull mysql
```

# Step 2 Preparing CLI
```
brew install redis
```

# Step 3  Start project
1. Start Docker Image for Redis with the following command
```
 docker compose up mysql redis 
```

2. Start Go Service 
```
go run main.go
```


# Step 4  Call Api with PostMan
- Import collection postman from go_ktb_test/postman/ktb.postman_collection.json to postman 
- Create Product -> Using for create product
- Get All Product -> Using for get all product that be created
- Get Product By Id -> Using for get product by id
- Update Product By Id -> Using for update product by id
- Delete Product By Id -> Using for delete product by id
- Get Pokemon List (Call another REST API) ->  -> Using for getting pokemon data from another rest api




