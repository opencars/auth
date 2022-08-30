package http

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/opencars/httputil"

	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/auth/pkg/eventapi"
)

type AuthHandler struct {
	publisher eventapi.Publisher
	store     domain.Store
}

func (h *AuthHandler) handleAuth() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ipAddr := strings.Split(r.RemoteAddr, ",")[0]

		secret := h.apiKey(r)

		auth := model.Authorization{
			Status: model.AuthStatusSucceed,
			IP:     ipAddr,
			Time:   time.Now().UTC(),
			Token: model.Token{
				ID: secret,
			},
		}

		if secret == "" {
			return h.result(&auth, &ErrInvalidToken)
		}

		token, err := h.store.Token().FindBySecret(r.Context(), secret)
		if errors.Is(err, model.ErrTokenNotFound) {
			return h.result(&auth, &ErrInvalidToken)
		}

		if err != nil {
			return err
		}

		auth.Token = *token
		if !token.Enabled {
			return h.result(&auth, &ErrTokenRevoked)
		}

		item, err := h.store.Blacklist().FindByIPv4(auth.IP)
		if err == nil && item.Enabled {
			return h.result(&auth, &ErrAccessDenied)
		}

		if err != nil && !errors.Is(err, model.ErrBlacklistRecordNotFound) {
			return err
		}

		w.Header().Set(HeaderTokenID, token.ID)
		w.Header().Set(HeaderTokenName, token.Name)
		w.Header().Set(HeaderUserID, token.UserID)
		w.WriteHeader(http.StatusOK)

		return h.result(&auth, nil)
	}
}

func (h *AuthHandler) result(auth *model.Authorization, httpErr *httputil.Error) error {
	if httpErr != nil {
		auth.Status = model.AuthStatusFailed
		auth.Error = new(string)
		*auth.Error = (*httpErr).Error()
	}

	event, err := eventapi.NewEvent(eventapi.EventAuthorizationKind, &auth)
	if err != nil {
		return err
	}

	if err := h.publisher.Publish(event); err != nil {
		return err
	}

	if httpErr != nil {
		return *httpErr
	}

	return nil
}

func (h *AuthHandler) apiKey(r *http.Request) string {
	if v := r.Header.Get("X-Api-Key"); v != "" {
		return v
	} else if v := r.Header.Get("Api-Key"); v != "" {
		return v
	}

	return ""
}
