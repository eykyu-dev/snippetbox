package main

import "snippetbox.eykyu.dev/internals/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
