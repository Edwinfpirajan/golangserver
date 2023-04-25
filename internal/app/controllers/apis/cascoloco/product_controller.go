package controlles

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	entity "github.com/CASCOLOCO/boreal.git/internal/interfaces/entity/APIs/cascoloco"
	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	response, err := http.Get("https://tienda.scsintercom.com/api/products?display=[id,price,name,reference,stock_availables[id,id_product_attribute]]&output_format=JSON&ws_key=S7UVTH5XIPYEPRWRHD5SUZNVZKT8SU1I")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var data entity.ProductResponse

	err = json.NewDecoder(response.Body).Decode(&data)
	fmt.Println(&data)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, data)
}

func GetProductOptionValues(c echo.Context) error {
	response2, err := http.Get("https://tienda.scsintercom.com/api/product_option_values?display=[id,name]&output_format=JSON&ws_key=S7UVTH5XIPYEPRWRHD5SUZNVZKT8SU1I")
	if err != nil {
		log.Fatal(err)
	}

	defer response2.Body.Close()

	var data2 entity.ProductOptionValuesResponse

	err = json.NewDecoder(response2.Body).Decode(&data2)
	fmt.Println(&data2)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, data2)
}

func GetCombinationsAndProductOptionValues(c echo.Context) error {

	urlProduct := "https://tienda.scsintercom.com/api/products?display=[id,price,name,reference,stock_availables[id,id_product_attribute]]&output_format=JSON&ws_key=S7UVTH5XIPYEPRWRHD5SUZNVZKT8SU1I"
	urlCombinations := "https://tienda.scsintercom.com/api/combinations?display=[id,id_product,reference,product_option_values[id]]&output_format=JSON&ws_key=S7UVTH5XIPYEPRWRHD5SUZNVZKT8SU1I"
	// urlCombinations := "https://tienda.scsintercom.com/api/combinations?display=[id,id_product,quantity,reference,product_option_values[id]]&output_format=JSON&ws_key=S7UVTH5XIPYEPRWRHD5SUZNVZKT8SU1I"
	urlOptionsValues := "https://tienda.scsintercom.com/api/product_option_values?display=[id,name]&output_format=JSON&ws_key=S7UVTH5XIPYEPRWRHD5SUZNVZKT8SU1I"
	urlProductStock := "https://tienda.scsintercom.com/api/stock_availables?display=[id_product,id_product_attribute,quantity]&output_format=JSON&ws_key=S7UVTH5XIPYEPRWRHD5SUZNVZKT8SU1I"

	//urlProduct

	resp1, err := http.Get(urlProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp1.Body.Close()

	var products entity.ProductResponse
	err = json.NewDecoder(resp1.Body).Decode(&products)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// urlCombinations

	resp2, err := http.Get(urlCombinations)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp2.Body.Close()

	var productCombinations entity.ProductCombinationResponse
	err = json.NewDecoder(resp2.Body).Decode(&productCombinations)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	//urlOptionValues

	resp3, err := http.Get(urlOptionsValues)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp3.Body.Close()

	var productOptionValues entity.ProductOptionValuesResponse
	err = json.NewDecoder(resp3.Body).Decode(&productOptionValues)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp4, err := http.Get(urlProductStock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var producStock entity.ProducStockResponse
	err = json.NewDecoder(resp4.Body).Decode(&producStock)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := entity.ProductAllData{
		Products:            products,
		ProductCombinations: productCombinations,
		ProductOptionValues: productOptionValues,
		ProductStock:        producStock,
	}

	return c.JSON(http.StatusOK, &response)
}

// response := struct {
// 	Products            entity.ProductResponse            `json:"products"`
// 	ProductCombinations entity.ProductCombinationResponse `json:"combinations"`
// 	ProductOptionValues entity.ProductOptionValuesResponse
// }{
// 	Products:            Products,
// 	ProductCombinations: ProductCombinations,
// 	ProductOptionValues: ProductOptionValues,
// }
