package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zaenalarifin12/test-ecommerce/internal/domain"
)

type productRepository struct {
	db *pgxpool.Pool
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *pgxpool.Pool) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) Create(ctx context.Context, product *domain.Product) (int64, error) {
	var id int64
	err := repo.db.QueryRow(ctx, "INSERT INTO products(product_name, price, quantity) VALUES($1, $2, $3) RETURNING product_id", product.ProductName, product.Price, product.Quantity).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *productRepository) Read(ctx context.Context, id int32) (*domain.Product, error) {
	var p domain.Product
	err := repo.db.QueryRow(ctx, "SELECT product_id, product_name, price, quantity FROM products WHERE product_id = $1", &id).Scan(&p.ProductID, &p.ProductName, &p.Price, &p.Quantity)
	if err != nil {
		return &domain.Product{}, err
	}
	return &p, nil
}

func (repo *productRepository) Update(ctx context.Context, product *domain.Product, id int64) error {
	_, err := repo.db.Exec(ctx, "UPDATE products SET product_name = $1, price = $2, quantity = $3 WHERE product_id = $4", product.ProductName, product.Price, product.Quantity, id)
	return err
}

func (repo *productRepository) Delete(ctx context.Context, id int64) error {
	_, err := repo.db.Exec(ctx, "DELETE FROM products WHERE product_id = $1", id)
	return err
}

func (repo *productRepository) List(ctx context.Context) ([]domain.Product, error) {
	rows, err := repo.db.Query(ctx, "SELECT product_id, product_name, price, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.Price, &p.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
