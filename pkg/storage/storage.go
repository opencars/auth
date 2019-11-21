package storage

// Adapter represnt abstract interface for communication of storage tools.
type Adapter interface {
	Token(id string) (*Token, error)
}

// Token represents full information about the auth-method of the request.
type Token struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
}

// Store helps to work with data by using abstact methods.
type Store struct {
	adapter Adapter
}

// New creates new store.
func New(adapter Adapter) *Store {
	return &Store{
		adapter: adapter,
	}
}

// Token returns full information about the auth method by uniqnue id.
func (s *Store) Token(id string) (*Token, error) {
	return s.adapter.Token(id)
}
