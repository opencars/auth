package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Settings is decoded configuration file.
type Settings struct {
	DB       Database `yaml:"database"`
	EventAPI EventAPI `yaml:"event_api"`
	Log      Log      `yaml:"log"`
	Server   Server   `yaml:"server"`
}

// Log represents settings for application logger.
type Log struct {
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

// Server represents settings for creating http server.
type Server struct {
	ShutdownTimeout Duration `yaml:"shutdown_timeout"`
	ReadTimeout     Duration `yaml:"read_timeout"`
	WriteTimeout    Duration `yaml:"write_timeout"`
	IdleTimeout     Duration `yaml:"idle_timeout"`
}

// Database contains configuration details for database.
type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

type EventAPI struct {
	Enabled bool   `yaml:"enabled"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

func (api *EventAPI) Address() string {
	return fmt.Sprintf("nats://%s:%d", api.Host, api.Port)
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	var config Settings

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
