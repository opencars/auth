package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/auth/pkg/logger"
	"github.com/opencars/httputil"
)

var (
	// ErrInvalidToken returned, if ApiKey/X-Api-Key is not valid.
	ErrInvalidToken = httputil.NewError(http.StatusUnauthorized, "auth.token.is_not_valid")

	// ErrTokenRevoked returned, if ApiKey/X-Api-Key was temporary disabled.
	ErrTokenRevoked = httputil.NewError(http.StatusUnauthorized, "auth.token.revoked")

	// ErrAccessDenied returned, if ip address is blacklisted.
	ErrAccessDenied = httputil.NewError(http.StatusForbidden, "auth.access_denied")
)

func handleErr(err error) error {
	if err != nil {
		logger.Errorf("error: %+v", err)
	}

	var e model.Error
	if errors.As(err, &e) {
		return httputil.NewError(http.StatusBadRequest, e.Error())
	}

	var vErr model.ValidationError
	if errors.As(err, &vErr) {
		errMessage := make([]string, 0)
		for k, vv := range vErr.Messages {
			for _, v := range vv {
				errMessage = append(errMessage, fmt.Sprintf("%s.%s", k, v))
			}
		}

		return httputil.NewError(http.StatusBadRequest, errMessage...)
	}

	return err
}
