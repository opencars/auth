package http

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/httputil"
)

var (
	ErrUnauthorized = httputil.NewError(http.StatusUnauthorized, "user.not_authorized")
)

func SessionCheckerMiddleware(checker domain.SessionChecker) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return httputil.Handler(func(w http.ResponseWriter, r *http.Request) error {
			cookie := r.Header.Get("Cookie")
			sessionToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)

			user, err := checker.CheckSession(r.Context(), sessionToken, cookie)
			if err != nil {
				return ErrUnauthorized
			}

			ctx := WithUserID(r.Context(), user.ID)
			next.ServeHTTP(w, r.WithContext(ctx))

			return nil
		})
	}
}
