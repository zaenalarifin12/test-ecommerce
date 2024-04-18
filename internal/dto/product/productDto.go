package product

type ProductRequest struct {
	ProductName string `json:"product_name" `
	Price       int64  `json:"price" `
	Quantity    int64  `json:"quantity" `
}

type ProductResponse struct {
	ProductID   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}
