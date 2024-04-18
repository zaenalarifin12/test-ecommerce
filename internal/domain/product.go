package domain

import (
	"context"
	"github.com/zaenalarifin12/test-ecommerce/internal/dto/product"
)

type Product struct {
	ProductID   int64  `db:"product_id"`
	ProductName string `db:"product_name"`
	Price       int64  `db:"price"`
	Quantity    int64  `db:"quantity"`
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) (int64, error)
	Read(ctx context.Context, id int32) (*Product, error)
	Update(ctx context.Context, product *Product, id int64) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]Product, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *product.ProductRequest) (int64, error)
	GetProduct(ctx context.Context, id int32) (*product.ProductResponse, error)
	UpdateProduct(ctx context.Context, product *product.ProductRequest, id int64) error
	DeleteProduct(ctx context.Context, id int64) error
	ListProduct(ctx context.Context) ([]product.ProductResponse, error)
}
