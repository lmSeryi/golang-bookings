package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lmSeryi/bookings/internal/config"
	"github.com/lmSeryi/bookings/internal/forms"
	"github.com/lmSeryi/bookings/internal/modules"
	"github.com/lmSeryi/bookings/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Handler for the Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIp", remoteIp)

	render.RenderTemplate(w, r, "home.page.tmpl", &modules.TemplateData{
		StringMap: map[string]string{"test": "Hello, World."},
	})
}

// About Handler for the About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again."

	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")
	stringMap["remoteIp"] = remoteIp

	// send the data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &modules.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation Handler for the Reservation page
func (m Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &modules.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation handles the posting of a reservation form
func (m Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(r.Form)
	reservation := &modules.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	// validate the form
	form := forms.New(r.PostForm)

	form.Has("first_name", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &modules.TemplateData{
			Form: form,
			Data: data,
		})
	}

}

// Majors Handler for the Majors page
func (m Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &modules.TemplateData{})
}

// Generals Handler for the Generals page
func (m Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &modules.TemplateData{})
}

// Availability Handler for the Availability page
func (m Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &modules.TemplateData{})
}

// PostAvailability Handler for the PostAvailability page
func (m Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start: %s, end: %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson returns a JSON response of the availability
func (m Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact Handler for the Contact page
func (m Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &modules.TemplateData{})
}
