package handlers

import (
	"net/http"

	"github.com/tikimcrzx/bookings/cmd/pkg/config"
	"github.com/tikimcrzx/bookings/cmd/pkg/models"
	"github.com/tikimcrzx/bookings/cmd/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderHTMLTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

// Test is the test page handler
func (m *Repository) Test(rw http.ResponseWriter, r *http.Request) {
	render.RenderHTMLTemplate(rw, "test.page.tmpl", &models.TemplateData{})
}

// SingleRoom is the single-room page handler
func (m *Repository) SingleRoom(rw http.ResponseWriter, r *http.Request) {
	render.RenderHTMLTemplate(rw, "single-room.page.tmpl", &models.TemplateData{})
}

// DeluxeRoom is the single-room page handler
func (m *Repository) DeluxeRoom(rw http.ResponseWriter, r *http.Request) {
	render.RenderHTMLTemplate(rw, "deluxe-room.page.tmpl", &models.TemplateData{})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(rw http.ResponseWriter, r *http.Request) {
	render.RenderHTMLTemplate(rw, "reservation.page.tmpl", &models.TemplateData{})
}

// MakeReservation is the make reservation page handler
func (m *Repository) MakeReservation(rw http.ResponseWriter, r *http.Request) {
	render.RenderHTMLTemplate(rw, "make-reservation.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Cossette"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderHTMLTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
