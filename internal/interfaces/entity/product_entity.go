package entity

type ProductResponse struct {
	Products []struct {
		ID           int    `json:"id"`
		Reference    string `json:"reference"`
		Price        string `json:"price"`
		Name         string `json:"name"`
		Associations struct {
			StockAvailables []struct {
				ID                 string `json:"id"`
				IDProductAttribute string `json:"id_product_attribute"`
			} `json:"stock_availables"`
		} `json:"associations"`
	} `json:"products"`
}

type ProductCombination struct {
	ID        int64  `json:"id"`
	IDProduct string `json:"id_product"`
	// Quantity     int64  `json:"quantity"`
	Reference    string `json:"reference"`
	Associations struct {
		ProductOptionValues []struct {
			ID string `json:"id"`
		} `json:"product_option_values"`
	} `json:"associations"`
}

type ProductCombinationResponse struct {
	Combinations []ProductCombination `json:"combinations"`
}

type ProductOptionValue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductOptionValuesResponse struct {
	ProductOptionValues []ProductOptionValue `json:"product_option_values"`
}

type ProducStockResponse struct {
	Stock_availables []struct {
		ID                 string `json:"id"` // IDStockAvailable
		IDProduct          string `json:"id_product"`
		IDProductAttribute string `json:"id_product_attribute"`
		Quantity           string `json:"quantity"`
	} `json:"stock_availables"`
}

type ProductAllData struct {
	Products            ProductResponse
	ProductCombinations ProductCombinationResponse
	ProductOptionValues ProductOptionValuesResponse
	ProductStock        ProducStockResponse
}
