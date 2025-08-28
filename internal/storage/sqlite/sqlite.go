package sqlite

import (
	"database/sql"

	"github.com/mubinkg/foodi-exam/internal/config"
)

type SQlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*SQlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
}
