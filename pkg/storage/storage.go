package storage

type Database interface {
	Token(id string) (*Token, error)
}

type Token struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
}

type Store struct {
	db Database
}

func New(db Database) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Token(id string) (*Token, error) {
	return s.db.Token(id)
}
