package types

type Product struct {
	Id    int64   `json:"id"`
	Title string  `json:"title" validate:"required"`
	Body  string  `json:"body" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}
