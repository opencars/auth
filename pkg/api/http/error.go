package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/httputil"
)

func handleErr(err error) error {
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
