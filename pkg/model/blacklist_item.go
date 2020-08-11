package model

type BlackListItem struct {
	IPv4    string `json:"ipv4" db:"ipv4"`
	Enabled bool   `json:"enabled" db:"enabled"`
}
