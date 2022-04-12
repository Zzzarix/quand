package app

import (
	"context"
	"net/http"
	rep "quand/app/repository"
	"quand/pkg/parser"
	"time"
)

type App struct {
	server *http.Server
}

type Config struct {
	Port     string `env:"PORT" json:"port"`
	DbUsers  rep.Db `json:"dbUsers"`
	DbQuests rep.Db `json:"dbQuests"`
	Parsers  struct {
		Reddit parser.RedditParser `json:"reddit"`
	} `json:"parsers"`
	Api struct {
		Version string `json:"version"`
	} `json:"api"`
}

func (a *App) Run(config Config) error {
	handler := &Handler{Config: config}
	rep.Init(config.DbUsers, config.DbQuests)

	a.server = &http.Server{
		Addr:           ":" + config.Port,
		Handler:        handler.Init(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go timelyParsing(config.Parsers.Reddit)
	go timelyTodayQuestion()

	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
