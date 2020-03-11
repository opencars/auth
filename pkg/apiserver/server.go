package apiserver

import (
	"log"
	"net/http"
	"time"

	"github.com/opencars/auth/pkg/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/opencars/auth/pkg/eventapi"
	"github.com/opencars/auth/pkg/store"
)

type server struct {
	router    *mux.Router
	publisher eventapi.Publisher
	store     store.Store
}

func newServer(store store.Store, publisher eventapi.Publisher) *server {
	server := &server{
		router:    mux.NewRouter(),
		publisher: publisher,
		store:     store,
	}

	server.configureRoutes()

	return server
}

func (s *server) configureRoutes() {
	s.router.Handle("/api/v1/auth", s.handleAuth()).Methods("GET", "OPTIONS")
}

func (s *server) handleAuth() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var id string
		if apiKey := r.Header.Get("X-Api-Key"); apiKey != "" {
			id = apiKey
		} else {
			id = r.Header.Get("Api-Key")
		}

		auth := model.Authorization{
			Status: "succeed",
			IP:     r.RemoteAddr,
			Time:   time.Now().UTC(),
			Token: model.Token{
				ID: id,
			},
		}

		if id == "" {
			return s.result(&auth, &InvalidToken)
		}

		token, err := s.store.Token().FindByID(id)
		if err == store.ErrRecordNotFound {
			return s.result(&auth, &InvalidToken)
		}

		if err != nil {
			return err
		}

		auth.Token = *token
		if !token.Enabled {
			return s.result(&auth, &TokenRevoked)
		}

		log.Printf("%s - %s - %v\n", token.ID, token.Name, token.Enabled)

		w.Header().Set("X-Auth-Id", token.ID)
		w.Header().Set("X-Auth-Name", token.Name)
		w.WriteHeader(http.StatusOK)

		return s.result(&auth, nil)
	}
}

func (s *server) result(auth *model.Authorization, authErr *Error) error {
	if authErr != nil {
		auth.Status = "failed"
		auth.Error = &authErr.Message
	}

	event, err := eventapi.NewEvent(eventapi.EventAuthorizationKind, &auth)
	if err != nil {
		return err
	}

	if err := s.publisher.Publish(event); err != nil {
		log.Println(err)
	}

	if authErr != nil {
		return *authErr
	}

	return nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Api-Key", "Api-Key"})
	cors := handlers.CORS(origins, methods, headers)(s.router)

	cors.ServeHTTP(w, r)
}
