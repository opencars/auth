package auth

import (
	"github.com/opencars/auth/pkg/storage"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	store *storage.Store
}

// All requests will go here first to see if client is authenticated
func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("Api-Key")

	// No auth token - respond unauthorized.
	if id == "" {
		log.Println("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// TODO: Check apiKey by selecting it from the storage.
	token, err := h.store.Token(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("X-Name", token.Name)
	w.WriteHeader(http.StatusOK)
}

func main() {
	store := storage.Store{}

	router := mux.NewRouter()
	router.Handle("/api/v1/auth", AuthHandler{})

	srv := http.Server{
		Addr:    ":8080",
		Handler: handlers.LoggingHandler(os.Stdout, router),
	}

	log.Println("Listening on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
