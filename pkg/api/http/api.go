package http

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"

	"github.com/opencars/auth/pkg/config"
	"github.com/opencars/auth/pkg/domain"
	"github.com/opencars/auth/pkg/eventapi"
	"github.com/opencars/auth/pkg/logger"
)

// Start starts the server with specified store.
func Start(ctx context.Context, addr string, conf *config.Settings, pub eventapi.Publisher, store domain.Store, svc domain.UserService, checker domain.SessionChecker) error {
	s := newServer(pub, store, svc, checker, &conf.Kratos)

	srv := http.Server{
		Addr:           addr,
		Handler:        handlers.CustomLoggingHandler(os.Stdout, handlers.ProxyHeaders(s), logFormatter),
		ReadTimeout:    conf.Server.ReadTimeout.Duration,
		WriteTimeout:   conf.Server.WriteTimeout.Duration,
		IdleTimeout:    conf.Server.IdleTimeout.Duration,
		MaxHeaderBytes: 1 << 20,
	}

	errs := make(chan error)
	go func() {
		errs <- srv.ListenAndServe()
	}()

	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		ctxShutDown, cancel := context.WithTimeout(context.Background(), conf.Server.ShutdownTimeout.Duration)
		defer cancel()

		err := srv.Shutdown(ctxShutDown)
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	}
}

func logFormatter(_ io.Writer, pp handlers.LogFormatterParams) {
	logger.WithFields(logger.Fields{
		"method": pp.Request.Method,
		"path":   pp.URL.Path,
		"status": pp.StatusCode,
		"size":   pp.Size,
		"addr":   strings.Split(pp.Request.RemoteAddr, ",")[0],
	}).Infof("http")
}
