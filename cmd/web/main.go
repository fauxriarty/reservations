package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fauxriarty/reservations/internal/config"
	"github.com/fauxriarty/reservations/internal/handlers"
	"github.com/fauxriarty/reservations/internal/models"
	"github.com/fauxriarty/reservations/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	// what the session is going to store
	gob.Register(models.Reservation{})

	app.InProduction = false // set to true in production

	//sessions is like the sharedpref of websites,
	//stores data that needs to persist during browsing multiple pages of the website
	session = scs.New()
	session.Lifetime = 24 * time.Hour // session created will last for 24 hours
	// by default it stores the sessions in cookies

	session.Cookie.Persist = true
	// if true, the session will persist after the browser is closed

	session.Cookie.SameSite = http.SameSiteLaxMode
	//attribute that tells the browser to only send the cookie if
	// the request originated from the same site

	session.Cookie.Secure = app.InProduction // its a bool declared in appconfig
	// if true, the cookie will only be sent over https not http

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	// its false bec we're in development mode since we want to see realtime changes not the cached version
	app.UseCache = false // set to true in production for efficiency

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s ", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	_ = srv.ListenAndServe()

}
