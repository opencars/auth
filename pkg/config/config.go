package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// Settings is decoded configuration file.
type Settings struct {
	DB       Database `toml:"database"`
	EventAPI EventAPI `toml:"event_api"`
}

// Database contains configuration details for database.
type Database struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"username"`
	Password string `toml:"password"`
	Name     string `toml:"database"`
}

type EventAPI struct {
	Enabled bool   `toml:"enabled"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
}

func (api *EventAPI) Address() string {
	return fmt.Sprintf("nats://%s:%d", api.Host, api.Port)
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	config := new(Settings)
	if _, err := toml.DecodeFile(path, config); err != nil {
		log.Fatal(err)
	}

	return config, nil
}
