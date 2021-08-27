package handlers

import (
	"fmt"
	"github.com/mrcoggsworth/webapp/pkg/config"
	"github.com/mrcoggsworth/webapp/pkg/utils"
	"log"
	"net/http"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the type
type Repository struct {
	App config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: *a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// HomeHandler is a function to run when route hits /
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.page.tmpl")
	log.Printf(fmt.Sprintf("The %v page was visited...", r.URL.Path))
}

// AboutHandler is a function to run when route hits /about
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "about.page.tmpl")
	log.Printf(fmt.Sprintf("The %v page was visited...", r.URL.Path))
}
