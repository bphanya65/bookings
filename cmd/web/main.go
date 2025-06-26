package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/bphanya65/bookings/pkg/config"
	"github.com/bphanya65/bookings/pkg/handlers"
	"github.com/bphanya65/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // in production this should be true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Listening on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
