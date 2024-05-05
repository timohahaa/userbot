package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/timohahaa/postgres"
	"github.com/timohahaa/userbot/config"
	v1 "github.com/timohahaa/userbot/internal/controllers/http/v1"
	"github.com/timohahaa/userbot/internal/repository"
	"github.com/timohahaa/userbot/internal/service"
	"github.com/timohahaa/userbot/pkg/httpserver"
	"github.com/timohahaa/userbot/pkg/logger"
	"github.com/timohahaa/userbot/pkg/validator"
)

func Run(configFilePath string) {
	cfg, err := config.New(configFilePath)
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	mainLogger := logger.New("internal.log", cfg.Server.LogPath)
	httpLogger := logger.New("requests.log", cfg.Server.LogPath)

	// database
	mainLogger.Info("initializing postgres connection...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxConnPoolSize(cfg.PG.ConnPoolSize))
	if err != nil {
		mainLogger.WithFields(logrus.Fields{"error": err}).Fatal("error connecting to postgres")
	}

	// транспортный слой
	mainLogger.Info("initializing repositories...")
	telegramRepo := repository.NewTelegramRepo(pg)

	// слой БЛ
	mainLogger.Info("initializing services...")
	telegramService := service.NewTelegramService(
		telegramRepo,
		cfg.Telegram.ApiId,
		cfg.Telegram.ApiHash,
		cfg.Telegram.Bot.CommentFrequency,
	)

	// handlers and routes
	mainLogger.Info("initializing handlers and routes...")
	handler := v1.NewRouter(telegramService, httpLogger)
	handler.Validator = validator.New()

	mainLogger.Infof("starting http server...")
	server := httpserver.New(handler, httpserver.Port(cfg.Server.Port))

	// gracefull shutdown
	mainLogger.Info("configuring gracefull shutdown...")
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	mainLogger.WithFields(logrus.Fields{"port": cfg.Server.Port}).Info("server started!")

	mainLogger.Info("shutting down...")
	err = server.Shutdown()
	if err != nil {
		mainLogger.WithFields(logrus.Fields{"error": err}).Fatal("error shutting down the server")
	}
}
