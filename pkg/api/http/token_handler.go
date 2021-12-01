package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/auth/pkg/domain/command"
	"github.com/opencars/auth/pkg/domain/query"
	"github.com/opencars/auth/pkg/handler"
)

type tokenHandler struct {
	svc domain.UserService
}

func (h *tokenHandler) Create() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID := UserIDFromContext(r.Context())

		var req CreateTokenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}

		c := command.CreateToken{
			UserID: userID,
			Name:   req.Name,
		}

		res, err := h.svc.CreateToken(r.Context(), &c)
		if err != nil {
			return handleErr(err)
		}

		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(res)
	}
}

func (h *tokenHandler) Delete() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID := UserIDFromContext(r.Context())
		tokenID := mux.Vars(r)["token_id"]

		c := command.DeleteToken{
			UserID:  userID,
			TokenID: tokenID,
		}

		err := h.svc.DeleteToken(r.Context(), &c)
		if err != nil {
			return handleErr(err)
		}

		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

func (h *tokenHandler) List() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID := UserIDFromContext(r.Context())

		var req CreateTokenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}

		q := query.ListTokens{
			UserID: userID,
			Limit:  r.URL.Query().Get("limit"),
			Offset: r.URL.Query().Get("offset"),
		}

		res, err := h.svc.ListTokens(r.Context(), &q)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(res)
	}
}

func (h *tokenHandler) Reset() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID := UserIDFromContext(r.Context())
		tokenID := mux.Vars(r)["token_id"]

		c := command.ResetToken{
			UserID:  userID,
			TokenID: tokenID,
		}

		token, err := h.svc.ResetToken(r.Context(), &c)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(token)
	}
}

func (h *tokenHandler) Details() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		userID := UserIDFromContext(r.Context())
		tokenID := mux.Vars(r)["token_id"]

		c := query.TokenDetails{
			UserID:  userID,
			TokenID: tokenID,
		}

		token, err := h.svc.TokenDetails(r.Context(), &c)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(token)
	}
}
