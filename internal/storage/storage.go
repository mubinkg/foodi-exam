package storage

import "github.com/mubinkg/foodi-exam/internal/types"

type Storage interface {
	CreateProduct(title string, body string, price float64) (int64, error)
	GetProductById(id int64) (types.Product, error)
	GetAllProducts() ([]types.Product, error)
	UpdateProduct(id int64, title string, body string, price float64) error
	SearchProducts(query string, sort string) ([]types.Product, error)
}
