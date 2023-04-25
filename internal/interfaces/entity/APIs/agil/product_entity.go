package entity

type BodyProductPost struct {
	Company   string `json:"Empresa"`
	Code      string `json:"Codigo"`
	Store     string `json:"Bodega"`
	PriceList string `json:"ListaDePrecio"`
	Init      int    `json:"Inicio"`
	Limit     int    `json:"Limite"`
}

type ProductList struct {
	Size    int       `json:"size"`
	Start   int       `json:"start"`
	Limit   int       `json:"limit"`
	Results []Product `json:"results"`
}

type Product struct {
	SizeColor    string      `json:"size_color"`
	ID           int         `json:"_id"`
	Code         string      `json:"code"`
	Barcode      string      `json:"barcode"`
	Name         string      `json:"name"`
	Brand        string      `json:"brand"`
	Presentation string      `json:"presentation"`
	Price        []PriceInfo `json:"price"`
	Inventory    []Inventory `json:"inventory"`
	Grupo        string      `json:"grupo"`
	SubGrupo     string      `json:"sub_grupo"`
}

type PriceInfo struct {
	PriceCode    string  `json:"price_code"`
	Price        float64 `json:"price"`
	Presentation string  `json:"presentation"`
}

type Inventory struct {
	WineCellar string  `json:"wineCellar"`
	Inventory  float64 `json:"inventory"`
}
