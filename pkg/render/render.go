package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/lmSeryi/bookings/pkg/config"
	"github.com/lmSeryi/bookings/pkg/modules"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func AddDefaultData(td *modules.TemplateData) *modules.TemplateData {
	return td
}

// NewTemplates sets the config for the tempalte package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate helps to render a html template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *modules.TemplateData) {
	// get the template cache from the context
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// RenderTemplateWithData helps to render a html template with data
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
