package http

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/gorilla/handlers"

	"github.com/opencars/auth/pkg/config"
	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/auth/pkg/eventapi"
	"github.com/opencars/auth/pkg/version"
	"github.com/opencars/httputil"
)

type server struct {
	router    http.Handler
	publisher eventapi.Publisher
	store     domain.Store
}

func newServer(pub eventapi.Publisher, store domain.Store, svc domain.UserService, checker domain.SessionChecker, conf *config.Kratos) *server {
	s := server{
		router:    configureRouter(pub, store, svc, checker),
		publisher: pub,
		store:     store,
	}

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

func (*server) Version() httputil.Handler {
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
