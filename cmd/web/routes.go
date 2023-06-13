package main

import (
	"log"
	"net/http"

	"github.com/SubhaagChowdhury/project/pkg/config"
	"github.com/SubhaagChowdhury/project/pkg/handlers"
	"github.com/gorilla/mux"
)

func routes(app *config.AppConfig) http.Handler {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Fatal(e)
		}
	}()

	resp := mux.NewRouter()
	resp.Use(NoSurf)
	resp.Use(SessionLoad)
	resp.Handle("/", http.HandlerFunc(handlers.Repo.HomePage))
	resp.Handle("/about", http.HandlerFunc(handlers.Repo.About))

	return resp
}
