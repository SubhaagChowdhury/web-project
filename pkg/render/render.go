package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/SubhaagChowdhury/project/pkg/config"
	"github.com/SubhaagChowdhury/project/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get template cache from app config
	tmp, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data)
	tmp.Execute(buf, data)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	// get all files named *_page.html
	pages, err := filepath.Glob("./templates/*_page.html")
	if err != nil {
		return tmplCache, err
	}
	//range through all files ending with *_page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		matches, err := filepath.Glob("./templates/*_layout.html")
		if err != nil {
			return tmplCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*_layout.html")
			if err != nil {
				return tmplCache, err
			}
		}

		//update cache
		tmplCache[name] = ts
	}

	return tmplCache, nil
}
