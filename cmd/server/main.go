package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"

	"github.com/opencars/auth/pkg/apiserver"
	"github.com/opencars/auth/pkg/config"
	"github.com/opencars/auth/pkg/storage/postgres"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	// Get configuration.
	conf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Register postgres adapter.
	storage, err := postgres.New(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(":8080", storage); err != nil {
		log.Fatal(err)
	}
}
