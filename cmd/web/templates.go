package main

import (
	"html/template"
	"path/filepath"
	"snippetbox/internal/models"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/layouts/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.html",
			"./ui/html/nav.html",
			page,
		}
		tmp, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		cache[name] = tmp
	}
	return cache, nil
}
