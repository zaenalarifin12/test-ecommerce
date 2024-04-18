package service

import (
	"context"
	"github.com/zaenalarifin12/test-ecommerce/internal/domain"
	"github.com/zaenalarifin12/test-ecommerce/internal/dto/product"
)

type productService struct {
	productRepository domain.ProductRepository
}

func NewProduct(productRepository domain.ProductRepository) domain.ProductService {
	return &productService{productRepository: productRepository}
}

func (p *productService) CreateProduct(ctx context.Context, req *product.ProductRequest) (int64, error) {
	// Converting ProductRequest to domain.Product
	newProduct := &domain.Product{
		ProductName: req.ProductName,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	// Call repository method to create the product
	id, err := p.productRepository.Create(ctx, newProduct)
	return id, err
}

func (p *productService) GetProduct(ctx context.Context, id int32) (*product.ProductResponse, error) {
	// Call repository method to get the product by ID
	retrievedProduct, err := p.productRepository.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	// Convert domain.Product to ProductResponse
	response := &product.ProductResponse{
		ProductID:   retrievedProduct.ProductID,
		ProductName: retrievedProduct.ProductName,
		Price:       retrievedProduct.Price,
		Quantity:    retrievedProduct.Quantity,
	}

	return response, nil
}

func (p *productService) UpdateProduct(ctx context.Context, req *product.ProductRequest, id int64) error {
	// Converting ProductRequest to domain.Product
	updatedProduct := &domain.Product{
		ProductName: req.ProductName,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	// Call repository method to update the product
	err := p.productRepository.Update(ctx, updatedProduct, id)
	return err
}

func (p *productService) DeleteProduct(ctx context.Context, id int64) error {
	// Call repository method to delete the product by ID
	err := p.productRepository.Delete(ctx, id)
	return err
}

func (p *productService) ListProduct(ctx context.Context) ([]product.ProductResponse, error) {
	// Call repository method to list all products
	products, err := p.productRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	// Convert domain.Product to ProductResponse for each product
	var responses []product.ProductResponse

	for _, productDetail := range products {
		responses = append(responses,
			product.ProductResponse{
				ProductID:   productDetail.ProductID,
				ProductName: productDetail.ProductName,
				Price:       productDetail.Price,
				Quantity:    productDetail.Quantity,
			},
		)
	}

	return responses, nil
}
