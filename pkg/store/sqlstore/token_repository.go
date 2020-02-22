package sqlstore

import (
	"database/sql"

	"github.com/opencars/auth/pkg/model"
	"github.com/opencars/auth/pkg/store"
)

// TokenRepository is responsible for tokens.
type TokenRepository struct {
	store *Store
}

// FindByID returns full information about the auth method by uniqnue id.
func (r *TokenRepository) FindByID(id string) (*model.Token, error) {
	var token model.Token

	err := r.store.db.Get(&token, `SELECT id, name, enabled FROM tokens WHERE id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}
