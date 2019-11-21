package auth

import (
	"github.com/opencars/auth/pkg/storage"
	"log"
	"net/http"
)

type Handler struct {
	store *storage.Store
}

func NewHandler(store *storage.Store) *Handler {
	return &Handler{
		store: store,
	}
}

// All requests will go here first to see if client is authenticated
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
