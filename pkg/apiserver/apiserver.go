package apiserver

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/opencars/auth/pkg/storage"
)

// Start ...
func Start(addr string, store storage.Adapter) error {
	server := newServer(addr, store)

	srv := http.Server{
		Addr:    addr,
		Handler: handlers.LoggingHandler(os.Stdout, handlers.ProxyHeaders(server)),
	}

	log.Printf("Listening on port %s...\n", addr)
	return srv.ListenAndServe()
}
