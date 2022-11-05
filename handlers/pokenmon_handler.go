package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type pokemon struct {
	Count    int64           `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Results  []pokemonResult `json:"results"`
}

type pokemonResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetPokeMonList(c echo.Context) error {
	//Create New Request GET Api to open source api
	req, err := http.NewRequest(http.MethodGet, "https://pokeapi.co/api/v2/pokemon", nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 10 * time.Second}

	//Call Api by using req
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Read body data from request api
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var responsePokemon pokemon
	//Unmarshal body data to pokemon struct
	json.Unmarshal(body, &responsePokemon)

	return c.JSON(http.StatusOK, responsePokemon.Results)
}
