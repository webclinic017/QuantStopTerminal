package webserver

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/quantstop/quantstopterminal/internal/assets"
	"github.com/quantstop/quantstopterminal/internal/config"
	"github.com/quantstop/quantstopterminal/pkg/logger"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func StartHttpServer(config *config.Config) {

	var (
		err               error
		nodeContextCancel context.CancelFunc
	)

	logger.Infoln(logger.WebserverLogger, "Starting web server ...")

	/*
	 * Start the Go application
	 */
	httpServer := echo.New()
	httpServer.Use(middleware.CORS())

	httpServer.GET("/*", echo.WrapHandler(http.FileServer(assets.Assets)))
	httpServer.GET("/api/version", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, config.GetVersion(false))
	})

	go func() {

		err = httpServer.Start("0.0.0.0:8080")

		if err != http.ErrServerClosed {
			logger.Errorf(logger.WebserverLogger, "Error starting web server %v.\n", err)
		} else {
			logger.Infoln(logger.WebserverLogger, "Shutting down web server ...")
		}
		logger.Infoln(logger.WebserverLogger, "Starting web server ... Success.")

	}()

	/*
	 * Setup shutdown handler
	 */
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	/*
	 * Start the Node client app (only for version "development")
	 */
	if config.IsDevelopment {
		_, nodeContextCancel = context.WithCancel(context.Background())

		go func() {
			logger.Debugf(logger.WebserverLogger, "Starting node development server ...")

			var cmd *exec.Cmd
			var err error

			cmd = exec.Command("npm", "run", "serve")
			cmd.Dir = "../../web"
			cmd.Stdout = os.Stdout

			if err = cmd.Start(); err != nil {
				logger.Errorf(logger.WebserverLogger, "Error starting node development server %v.\n", err)
			}

			// Wait for command to stop running, ie. node server is stopped
			_ = cmd.Wait()

			logger.Infoln(logger.WebserverLogger, "Shutting down node development server ...")

		}()
	}

	/*
	 * Wait for and stop both the Go and Node apps
	 */
	<-quit

	if config.IsDevelopment {
		nodeContextCancel()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = httpServer.Shutdown(ctx); err != nil {
		logger.Errorf(logger.WebserverLogger, "Error starting web server! %v.\n", err)
	}

	logger.Infoln(logger.WebserverLogger, "Application stopped")
}
