package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/DanielTitkov/weavle/internal/app"
	"github.com/DanielTitkov/weavle/internal/configs"
	"github.com/DanielTitkov/weavle/internal/ent"
	"github.com/DanielTitkov/weavle/internal/logger"
	"github.com/DanielTitkov/weavle/internal/prepare"
	"github.com/DanielTitkov/weavle/internal/repository"

	_ "github.com/lib/pq"
)

func main() {
	// init configs
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	log.Println("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	log.Println("loaded config")

	// init logger
	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()
	logger.Info("starting service", "")

	// init database
	var dbOptions []ent.Option

	if cfg.Env == "dev" {
		if cfg.Debug.LogDBQueries {
			dbOptions = append(dbOptions, ent.Debug())
		}
	}
	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI, dbOptions...)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	// migrations
	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	// init repository
	repo := repository.NewRepository(db, logger)

	// store := prepare.Store(cfg)

	a, err := app.New(cfg, logger, repo)
	if err != nil {
		logger.Fatal("failed to init app", err)
	}

	fmt.Println(a)

	// jobs, err := job.New(cfg, logger, a)
	// if err != nil {
	// 	logger.Fatal("failed to init jobs", err)
	// }
	// jobs.Run() // async

	// gothic.Store = store.Store
	// goth.UseProviders(
	// 	google.New(
	// 		cfg.Auth.Google.Client,   // client
	// 		cfg.Auth.Google.Secret,   // secret
	// 		cfg.Auth.Google.Callback, // callback url
	// 		"email", "profile",       // scopes
	// 	),
	// 	github.New(
	// 		cfg.Auth.Github.Client,
	// 		cfg.Auth.Github.Secret,
	// 		cfg.Auth.Github.Callback,
	// 		"email", "profile",
	// 	),
	// 	twitter.NewAuthenticate(
	// 		cfg.Auth.Twitter.Client,
	// 		cfg.Auth.Twitter.Secret,
	// 		cfg.Auth.Twitter.Callback,
	// 	),
	// )

	// h := handler.NewHandler(a, logger, "templates/")
	// r := prepare.Mux(cfg, store, h)

	// httpServer := prepare.Server(cfg, r)
	// httpServer.Addr = cfg.Server.GetAddress()
	// logger.Info("starting http server", cfg.Server.GetAddress())
	// log.Fatal(httpServer.ListenAndServe())
}
