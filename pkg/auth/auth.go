package auth

import (
	"net/http"

	"github.com/opencars/auth/pkg/storage"
)

// Handler is responsible for checking request authentication.
// Validates "Api-Key" header to have right credentials.
type Handler struct {
	store *storage.Store
}

// NewHandler creates new instance of Handler.
func NewHandler(store *storage.Store) *Handler {
	return &Handler{
		store: store,
	}
}

// ServeHTTP implements http.Handler method.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id string
	if apiKey := r.Header.Get("X-Api-Key"); apiKey != "" {
		id = apiKey
	} else {
		id = r.Header.Get("Api-Key")
	}

	// No auth token - respond unauthorized.
	if id == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := h.store.Token(id)
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
