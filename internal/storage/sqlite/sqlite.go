package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/mubinkg/foodi-exam/internal/config"
	"github.com/mubinkg/foodi-exam/internal/types"
	_ "modernc.org/sqlite"
)

type SQlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*SQlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (id INTEGER PRIMARY KEY, title TEXT, body TEXT, price REAL)`)
	
	if err != nil {
		return nil, err
	}
	return &SQlite{Db: db}, nil
}

func (s *SQlite) CreateProduct(title string, body string, price float64) (int64, error){
	stmt, err := s.Db.Prepare(`INSERT INTO products (title, body, price) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(title, body, price)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *SQlite) GetProductById(id int64) (types.Product, error) {
	stmnt, err := s.Db.Prepare(`SELECT id, title, body, price FROM products WHERE id = ?`)
	if err != nil {
		return types.Product{}, err
	}
	defer stmnt.Close()

	var product types.Product
	err = stmnt.QueryRow(id).Scan(&product.Id, &product.Title, &product.Body, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Product{}, fmt.Errorf("product not found with id %d", id)
		}
		return types.Product{}, err
	}
	return product, nil
}
