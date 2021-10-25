package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/cikupin/kcd2021_helmfile/internal/bootstrap"
	"github.com/cikupin/kcd2021_helmfile/internal/logic"
	"github.com/go-chi/chi/v5"
	"github.com/urfave/cli/v2"
)

var API = &cli.Command{
	Name:        "api",
	Description: "Run api for KCD 2021 demo app",
	Action: func(c *cli.Context) error {
		runAPI()
		return nil
	},
}

func runAPI() error {
	config := bootstrap.LoadConfig()

	db, err := bootstrap.NewMysqlDatabase(config.DBOptions)
	if err != nil {
		log.Printf("[ERROR] Fail initiating database: %s\n", err.Error())
		return err
	}

	cacheClient := bootstrap.NewCache(config.CacheOptions)

	logic := logic.NewLogic(db, cacheClient, config)

	// create router
	r := chi.NewRouter()
	r.Get("/", logic.GetFrozenFoods)
	r.Get("/health", logic.HealthCheck)

	// initialize server
	var srv http.Server
	srv.Addr = fmt.Sprintf("0.0.0.0:%d", config.AppOptions.AppPort)
	srv.Handler = r

	startServer(&srv)
	return nil
}

func startServer(srv *http.Server) {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("[ERROR] Error on HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Printf("[INFO] API service is serving at %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("[ERROR] Error on HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	log.Printf("[INFO] API service has been stopped. Bye....")
}
