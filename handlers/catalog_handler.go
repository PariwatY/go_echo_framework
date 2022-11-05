package handlers

import (
	"encoding/json"
	"fmt"
	"go_ktb_test/constant"
	model "go_ktb_test/model/product"
	"go_ktb_test/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type catalogHandler struct {
	catalogSrv services.CatalogService
}

func NewCatalogHandler(catalogSrv services.CatalogService) CatalogHandler {
	return catalogHandler{catalogSrv}
}

func (h catalogHandler) GetProducts(c echo.Context) error {
	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := echo.Map{
		"status":   "ok",
		"products": products,
	}

	return c.JSON(http.StatusOK, response)
}

func (h catalogHandler) GetProductById(c echo.Context) error {
	//Get id from url param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("not found")
		return err
	}

	//Call GetProductById to get product by id
	product, err := h.catalogSrv.GetProductById(id)
	//Check product not found then return 404 not found status
	if err == constant.ErrNotFound {
		response := echo.Map{
			"status": constant.ErrNotFound.Error(),
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := echo.Map{
		"status":   "ok",
		"products": product,
	}
	return c.JSON(http.StatusOK, response)
}

func (h catalogHandler) SaveProduct(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	} else {
		//Get name and quantity from json_map that send be request body
		name := json_map["name"]
		quantity := int(json_map["quantity"].(float64))

		productDataCreated, err := h.catalogSrv.SaveProduct(name.(string), quantity)
		//Check error duplicate then return 201 created status
		if err == constant.ErrDuplicate {
			response := echo.Map{
				"status": err.Error(),
			}
			return c.JSON(http.StatusCreated, response)
		}

		//Create created response if nothing error happened and return 200 ok status
		response := echo.Map{
			"status":   "created",
			"products": productDataCreated,
		}
		return c.JSON(http.StatusOK, response)
	}

}
func (h catalogHandler) UpdateProductById(c echo.Context) error {

	productRequest := new(model.ProductRequest)

	if err := c.Bind(productRequest); err != nil {
		return err
	}

	productData := model.ProductRequest{
		ID:       productRequest.ID,
		Name:     productRequest.Name,
		Quantity: productRequest.Quantity,
	}

	productDataUpdated, err := h.catalogSrv.UpdateProductById(productData)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, productDataUpdated)
}
func (h catalogHandler) DeleteProductById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = h.catalogSrv.DeleteProductById(id)
	if err != nil {
		return err
	}

	response := echo.Map{
		"status": "delete successed",
	}
	return c.JSON(http.StatusOK, response)
}
