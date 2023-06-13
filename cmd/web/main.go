package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SubhaagChowdhury/project/pkg/config"
	"github.com/SubhaagChowdhury/project/pkg/handlers"
	"github.com/SubhaagChowdhury/project/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannout create template cache")
	}
	app.TemplateCache = tempCache

	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Running application in port " + port)

	serve := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
