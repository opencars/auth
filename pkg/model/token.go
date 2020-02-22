package model

type Token struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
}
