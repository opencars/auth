package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/auth/pkg/auth"
	"github.com/opencars/auth/pkg/config"
	"github.com/opencars/auth/pkg/storage"
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
	db, err := postgres.New(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name)
	if err != nil {
		log.Fatal(err)
	}

	store := storage.New(db)
	router := mux.NewRouter()
	router.Handle("/api/v1/auth", auth.NewHandler(store))

	srv := http.Server{
		Addr:    ":8080",
		Handler: handlers.LoggingHandler(os.Stdout, router),
	}

	log.Println("Listening on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
