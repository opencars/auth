package http

func (s *server) configureRouter() {
	router := s.router.PathPrefix("/api/v1/").Subrouter()

	router.Handle("/public/version", s.Version()).Methods("GET", "OPTIONS")

	// GET /api/v1/private/auth/{...}.
	auth := router.PathPrefix("/private/auth/")
	auth.Handler(s.handleAuth()).Methods("GET", "POST", "OPTIONS")
}
