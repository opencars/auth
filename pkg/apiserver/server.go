package apiserver

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/auth/pkg/storage"
)

type server struct {
	router *mux.Router
	store  storage.Adapter
}

func newServer(addr string, store storage.Adapter) *server {
	server := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	server.configureRoutes()

	return server
}

func (s *server) configureRoutes() {
	s.router.Handle("/api/v1/auth", s.handleAuth()).Methods("GET", "OPTIONS")
}

func (s *server) handleAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		for k, v := range r.Header {
			fmt.Println(k, strings.Join(v, " "))
		}

		var id string
		if apiKey := r.Header.Get("X-Api-Key"); apiKey != "" {
			id = apiKey
		} else {
			id = r.Header.Get("Api-Key")
		}

		if id == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := s.store.Token(id)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Enabled {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("X-Auth-Id", token.ID)
		w.Header().Set("X-Auth-Name", token.Name)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Api-Key", "Api-Key"})
	cors := handlers.CORS(origins, methods, headers)(s.router)

	cors.ServeHTTP(w, r)
}
