package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/opencars/auth/pkg/api/http"
	"github.com/opencars/auth/pkg/config"
	"github.com/opencars/auth/pkg/domain/service"
	"github.com/opencars/auth/pkg/eventapi/natspub"
	"github.com/opencars/auth/pkg/kratos"
	"github.com/opencars/auth/pkg/logger"
	"github.com/opencars/auth/pkg/store/sqlstore"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()

	conf, err := config.New(configPath)
	if err != nil {
		logger.Fatalf("failed read config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	store, err := sqlstore.New(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.SSLMode)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	pub, err := natspub.New(conf.EventAPI.Address(), conf.EventAPI.Enabled)
	if err != nil {
		logger.Fatalf("nats: %v", err)
	}

	svc := service.NewUserService(store.Token())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-c
		cancel()
	}()

	client, err := kratos.NewClient(conf.Kratos.BaseURL)
	if err != nil {
		logger.Fatalf("kratos: %v", err)
	}

	addr := ":8080"
	logger.Infof("Listening on %s...", addr)
	if err := http.Start(ctx, addr, conf, pub, store, svc, client); err != nil {
		logger.Fatalf("http server failed: %v", err)
	}
}
