package postgres

import "github.com/opencars/auth/pkg/storage"

type Database struct {
	sqlx.DB
}

func (db *Database) Token(id string) (*storage.Token, error) {
	
}
