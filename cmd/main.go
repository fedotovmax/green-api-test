package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fedotovmax/green-api-test/internal/app"
	"github.com/fedotovmax/green-api-test/internal/config"
	"github.com/fedotovmax/green-api-test/pkg/logger"
)

func setupLooger(env config.AppEnv) *slog.Logger {
	if env == config.Development {
		return logger.NewHandler(slog.LevelDebug)
	}
	return logger.NewHandler(slog.LevelWarn)
}

func main() {
	appConfig, err := config.New()

	if err != nil {
		logger.GetFallback().Error(err.Error())
		os.Exit(1)
	}

	log := setupLooger(appConfig.Env)

	app, err := app.New(appConfig, log)

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	notifyCtx, cancelNotifyCtx := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancelNotifyCtx()

	startErrChan := app.Start()

	select {
	case err := <-startErrChan:
		log.Error("Error recivied when starting application", logger.Err(err))
		cancelNotifyCtx()
	case <-notifyCtx.Done():
		log.Info("OS signal recevied")
	}

	log.Info("Starting to shutdown all resources")

	shutdownCtx, cancelShutdownCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdownCtx()

	app.Stop(shutdownCtx)

}
