package handlers

import (
	"github.com/Reymartinmontes/bookings/pkg/config"
	"github.com/Reymartinmontes/bookings/pkg/models"
	"github.com/Reymartinmontes/bookings/pkg/render"
	"net/http"
)

// Home is the home hanlder
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remove_ip", remoteIP)

	render.RenderTemplates(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	//send the data to the template
	render.RenderTemplates(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
