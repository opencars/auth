package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/auth/pkg/storage"
)

type wrapper struct {
	db *sqlx.DB
}

// New creates new implementation of storage.Adapter based on PostgreSQL.
func New(host string, port int, user, password, dbname string) (storage.Adapter, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &wrapper{
		db: db,
	}, nil
}

// Token returns full information about the auth method by uniqnue id.
func (w *wrapper) Token(id string) (*storage.Token, error) {
	var token storage.Token
	if err := w.db.Get(&token, `SELECT id, name, enabled FROM tokens WHERE id = $1`, id); err != nil {
		return nil, err
	}

	return &token, nil
}
