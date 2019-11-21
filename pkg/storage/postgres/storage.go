package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"github.com/opencars/auth/pkg/storage"
)

type database struct {
	db *sqlx.DB
}

func New(host string, port int, user, password, dbname string) (storage.Database, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &database{
		db: db,
	}, nil
}

func (db *database) Token(id string) (*storage.Token, error) {
	var token storage.Token
	if err := db.db.Get(&token, `SELECT id, name, enabled FROM tokens WHERE id = $1`, id); err != nil {
		return nil, err
	}

	return &token, nil
}
