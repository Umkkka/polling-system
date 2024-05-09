package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"polling-system/internal/config"
	"polling-system/internal/service/poll"
	"polling-system/internal/transport/handler"
	"polling-system/internal/transport/repository"
)

func setupLogger(config *config.Config) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}

	logger.SetLevel(logLevel)

	return logger, nil
}

func startServer(ctx context.Context, handler http.Handler, config *config.Config, logger *logrus.Logger) error {
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%v", config.ListenPort),
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
		IdleTimeout:  config.IdleTimeout,
		Handler:      handler,
	}

	go func() {
		<-ctx.Done()

		logger.Infof("Shutting down server with timeout %s...", config.ShutdownTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Fatal("Server shutdown failed: ", err)
		}

	}()

	logger.Info("Serving requests on :", config.ListenPort)

	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

func Run(config config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := setupLogger(&config)
	if err != nil {
		return fmt.Errorf("failed setup logger: %w", err)
	}

	repo := repository.New()
	pollService := poll.New(repo)
	pollHandler := handler.New(pollService)

	engine := initRouter(&config)

	setupTechnicalRoutes(engine)
	setupApiRoutes(engine, pollHandler)

	engine.GET("/ws", func(c *gin.Context) {
		handleWebSocketConnection(c.Writer, c.Request)
	})

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigint
		cancel()
	}()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return startServer(ctx, engine, &config, logger)
	})

	if err := group.Wait(); err != nil && err != context.Canceled {
		return err
	}

	return nil
}
