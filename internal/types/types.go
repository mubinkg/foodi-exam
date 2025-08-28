package types

type Product struct {
	Id    int
	Title string  `validate:"required"`
	Body  string  `validate:"required"`
	Price float64 `validate:"required"`
}
