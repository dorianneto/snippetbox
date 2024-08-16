package main

import "snippetbox.dorianneto.com/internal/models"

type TemplateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
