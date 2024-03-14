package handlers

import (
	"net/http"

	"github.com/fauxriarty/reservations/pkg/config"
	"github.com/fauxriarty/reservations/pkg/models"
	"github.com/fauxriarty/reservations/pkg/render"
)

// repo is the repository used by the handlers
var Repo *Repository

// repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	strings := make(map[string]string)
	strings["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strings["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: strings,
	})
}
