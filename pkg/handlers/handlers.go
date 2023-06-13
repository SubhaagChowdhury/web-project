package handlers

import (
	"net/http"

	"github.com/SubhaagChowdhury/project/pkg/config"
	"github.com/SubhaagChowdhury/project/pkg/models"
	"github.com/SubhaagChowdhury/project/pkg/render"
)

// TemplateData holds data sent from handlers to templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home_page.html", &models.TemplateData{})
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["test"] = "Hello, again"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	strMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about_page.html", &models.TemplateData{
		StringMap: strMap})
}
