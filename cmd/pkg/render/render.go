package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tikimcrzx/bookings/cmd/pkg/config"
	"github.com/tikimcrzx/bookings/cmd/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the templates package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefault initialize templateData
func AddDefault(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// RenderTemplate renders template using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	templateName, ok := templateCache[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
		return
	}

	buf := new(bytes.Buffer)

	templateData = AddDefault(templateData)
	_ = templateName.Execute(buf, templateData)
	_, err := buf.WriteTo(rw)

	if err != nil {
		fmt.Println("Error writting template to browser", err)
		return
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
