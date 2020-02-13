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

}
