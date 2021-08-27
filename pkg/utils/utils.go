package utils

import (
	"bytes"
	"github.com/mrcoggsworth/webapp/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the new config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string){
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("[!] could not get template from template cache...")
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("Error writing template to browser: %v", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.tmpl")
	if err != nil {
		log.Fatalln(err)
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			log.Fatalln(err)
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			log.Fatalln(err)
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		if err != nil {
			log.Fatalln(err)
			return cache, err
		}
		cache[name] = ts
	}
	return cache, nil
}
