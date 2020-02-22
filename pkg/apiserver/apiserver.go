package apiserver

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/opencars/auth/pkg/store"
)

// Start begins listening for requests.
func Start(addr string, store store.Store) error {
	server := newServer(store)

	srv := http.Server{
		Addr:    addr,
		Handler: handlers.LoggingHandler(os.Stdout, handlers.ProxyHeaders(server)),
	}

	log.Printf("Listening on port %s...\n", addr)
	return srv.ListenAndServe()
}
