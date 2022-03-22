package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"golang-bookings/internal/config"
	"golang-bookings/internal/handlers"
	"golang-bookings/internal/helpers"
	"golang-bookings/internal/models"
	"golang-bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {

	err := run()

	if err != nil {
		log.Fatal(err)
	}

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

func run() error {
	gob.Register(models.Reservation{})

	app.InProduction = false

	app.Infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.Errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour              // Cookie expires after 24 hours
	session.Cookie.Persist = true                  // Cookie will be saved even if the browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode // Cookie will be sent even if the site is not served over HTTPS
	session.Cookie.Secure = false                  // Cookie encrypted
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
