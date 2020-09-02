package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"

	"github.com/opencars/auth/pkg/apiserver"
	"github.com/opencars/auth/pkg/config"
	"github.com/opencars/auth/pkg/eventapi/natspub"
	"github.com/opencars/auth/pkg/store/sqlstore"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()

	// Get configuration.
	conf, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	store, err := sqlstore.New(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.SSLMode)
	if err != nil {
		log.Fatal(err)
	}

	pub, err := natspub.New(conf.EventAPI.Address(), conf.EventAPI.Enabled)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(":8080", store, pub); err != nil {
		log.Fatal(err)
	}
}
