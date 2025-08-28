package storage

type Storage interface {
	CreateProduct(title string, body string, price float64) (int64, error)
}
