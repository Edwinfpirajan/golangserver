package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	entity "github.com/CASCOLOCO/boreal.git/internal/interfaces/entity/APIs/agil"
	"github.com/labstack/echo/v4"
)

func GetAllProductsManual(c echo.Context) error {

	var body entity.BodyProductPost
	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req, err := http.NewRequest("POST", "http://76.74.150.33:8080/products/all", bytes.NewBuffer(requestBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	// fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var products entity.ProductList
	err = json.Unmarshal(respBody, &products)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// fmt.Println("Productos:", products)

	return c.JSON(http.StatusOK, products)
}
