package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/lmSeryi/bookings/internal/config"
	"github.com/lmSeryi/bookings/internal/handlers"
	"github.com/lmSeryi/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour              // Cookie expires after 24 hours
	session.Cookie.Persist = true                  // Cookie will be saved even if the browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode // Cookie will be sent even if the site is not served over HTTPS
	session.Cookie.Secure = false                  // Cookie encrypted
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Server is listening on port " + portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
