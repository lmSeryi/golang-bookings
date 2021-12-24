package handlers

import (
	"net/http"

	"github.com/lmSeryi/bookings/pkg/config"
	"github.com/lmSeryi/bookings/pkg/modules"
	"github.com/lmSeryi/bookings/pkg/render"
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

	render.RenderTemplate(w, "home.page.tmpl", &modules.TemplateData{})
}

// About Handler for the About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, Again."

	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")
	stringMap["remoteIp"] = remoteIp

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &modules.TemplateData{
		StringMap: stringMap,
	})
}
