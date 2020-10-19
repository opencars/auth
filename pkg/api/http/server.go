package http

import (
	"encoding/json"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/opencars/auth/pkg/eventapi"
	"github.com/opencars/auth/pkg/handler"
	"github.com/opencars/auth/pkg/model"
	"github.com/opencars/auth/pkg/store"
	"github.com/opencars/auth/pkg/version"
)

type server struct {
	router    *mux.Router
	publisher eventapi.Publisher
	store     store.Store
}

func newServer(pub eventapi.Publisher, store store.Store) *server {
	s := server{
		router:    mux.NewRouter(),
		publisher: pub,
		store:     store,
	}

	s.configureRouter()

	return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Api-Key", "X-Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	compress := handlers.CompressHandler(cors)
	compress.ServeHTTP(w, r)
}

func (*server) Version() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		v := struct {
			Version string `json:"version"`
			Go      string `json:"go"`
		}{
			Version: version.Version,
			Go:      runtime.Version(),
		}

		return json.NewEncoder(w).Encode(v)
	}
}

func (s *server) handleAuth() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ipAddr := strings.Split(r.RemoteAddr, ",")[0]

		apiKey := s.apiKey(r)
		auth := model.Authorization{
			Status: model.AuthStatusSucceed,
			IP:     ipAddr,
			Time:   time.Now().UTC(),
			Token: model.Token{
				ID: apiKey,
			},
		}

		if apiKey == "" {
			return s.result(&auth, &handler.ErrInvalidToken)
		}

		token, err := s.store.Token().FindByID(apiKey)
		if err == store.ErrRecordNotFound {
			return s.result(&auth, &handler.ErrInvalidToken)
		}

		if err != nil {
			return err
		}

		auth.Token = *token
		if !token.Enabled {
			return s.result(&auth, &handler.ErrTokenRevoked)
		}

		item, err := s.store.Blacklist().FindByIPv4(auth.IP)
		if err == nil && item.Enabled {
			return s.result(&auth, &handler.ErrAccessDenied)
		}

		if err != nil && err != store.ErrRecordNotFound {
			return err
		}

		w.Header().Set("X-Auth-Id", token.ID)
		w.Header().Set("X-Auth-Name", token.Name)
		w.WriteHeader(http.StatusOK)

		return s.result(&auth, nil)
	}
}

func (s *server) result(auth *model.Authorization, httpErr *handler.Error) error {
	if httpErr != nil {
		auth.Status = model.AuthStatusFailed
		auth.Error = new(string)
		*auth.Error = (*httpErr).Error()
	}

	event, err := eventapi.NewEvent(eventapi.EventAuthorizationKind, &auth)
	if err != nil {
		return err
	}

	if err := s.publisher.Publish(event); err != nil {
		return err
	}

	if httpErr != nil {
		return *httpErr
	}

	return nil
}

func (s *server) apiKey(r *http.Request) string {
	if v := r.Header.Get("X-Api-Key"); v != "" {
		return v
	} else if v := r.Header.Get("Api-Key"); v != "" {
		return v
	}

	return ""
}
