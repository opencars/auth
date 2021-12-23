package sqlstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/auth/pkg/domain/query"
)

// TokenRepository is responsible for tokens.
type TokenRepository struct {
	store *Store
}

func (r *TokenRepository) Create(ctx context.Context, token *model.Token) error {
	_, err := r.store.db.NamedExecContext(ctx,
		`INSERT INTO tokens
		(
			id, user_id, secret, name, enabled, created_at
		) VALUES(
			:id, :user_id, :secret, :name, :enabled, :created_at
		)`, token,
	)

	return err
}

func (r *TokenRepository) Update(ctx context.Context, token *model.Token) error {
	_, err := r.store.db.NamedExecContext(ctx,
		`UPDATE tokens SET secret= :secret name=:name enabled=:enabled
		WHERE id =:id)`,
		token,
	)

	return err
}

func (r *TokenRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.store.db.NamedExecContext(ctx, `DELETE tokens WHERE id = $1`, id)

	return err
}

// FindByID returns full information about the auth method by uniqnue id.
func (r *TokenRepository) FindByID(ctx context.Context, id string) (*model.Token, error) {
	var token model.Token

	err := r.store.db.GetContext(ctx, &token,
		`SELECT id, user_id, secret, name, enabled, created_at, enabled, updated_at  FROM tokens WHERE id = $1`,
		id,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrTokenNotFound
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// FindBySecret returns full information about the auth method by uniqnue secret.
func (r *TokenRepository) FindBySecret(ctx context.Context, secret string) (*model.Token, error) {
	var token model.Token

	err := r.store.db.GetContext(ctx, &token,
		`SELECT id, user_id, secret, name, enabled, created_at, updated_at FROM tokens WHERE secret = $1`,
		secret,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrTokenNotFound
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *TokenRepository) List(ctx context.Context, q *query.ListTokens) ([]model.Token, error) {
	tokens := make([]model.Token, 0)

	err := r.store.db.Select(&tokens,
		`SELECT id, user_id, secret, name, enabled, created_at, updated_at
		FROM tokens
		ORDER BY created_at ASC
		LIMIT $1 OFFSET $2`,
		q.GetLimit(), q.GetOffset(),
	)

	if err != nil {
		return nil, err
	}

	return tokens, nil
}
