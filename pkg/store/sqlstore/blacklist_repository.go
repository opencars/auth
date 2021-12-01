package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/opencars/auth/pkg/domain/model"
)

// BlacklistRepository is responsible for black list.
type BlacklistRepository struct {
	store *Store
}

// FindByIPv4 returns full information about the blacklisted record method by unique id.
func (r *BlacklistRepository) FindByIPv4(ipv4 string) (*model.BlackListItem, error) {
	var item model.BlackListItem

	err := r.store.db.Get(&item, `SELECT ipv4, enabled FROM blacklist WHERE ipv4 = $1`, ipv4)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrBlacklistRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &item, nil
}
