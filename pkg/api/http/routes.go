package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/auth/pkg/eventapi"
)

func configureRouter(pub eventapi.Publisher, store domain.Store, svc domain.UserService, checker domain.SessionChecker) http.Handler {
	tokens := tokenHandler{svc: svc}
	tAuth := AuthHandler{publisher: pub, store: store}

	router := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()
	// TODO: Add version endpoint.
	// router.Handle("/public/version", s.Version()).Methods("GET", "OPTIONS")

	// GET /api/v1/private/auth/{...}.
	router.PathPrefix("/private/auth").Handler(tAuth.handleAuth()).Methods("GET", "POST", "OPTIONS")

	user := router.PathPrefix("/user").Subrouter()
	user.Use(
		SessionCheckerMiddleware(checker),
	)

	user.Handle("/tokens", tokens.Create()).Methods("POST")                 // POST   /api/v1/tokens.
	user.Handle("/tokens", tokens.List()).Methods("GET")                    // GET    /api/v1/tokens.
	user.Handle("/tokens/{token_id}", tokens.Details()).Methods("GET")      // GET    /api/v1/tokens/:id.
	user.Handle("/tokens/{token_id}", tokens.Delete()).Methods("DELETE")    // DELETE /api/v1/tokens/:id.
	user.Handle("/tokens/{token_id}/reset", tokens.Reset()).Methods("POST") // DELETE /api/v1/tokens/:id/reset.

	return router
}
