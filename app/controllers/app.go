package controllers

import (
	"github.com/revel/revel"
)

// App Controller
type App struct {
	GormController
}

// Index Views
func (c App) Index() revel.Result {
	page := map[string]interface{}{
		"title":       "The Official Site of Ngurajeka",
		"description": "Web Developer Ganteng ;)",
		"homepage":    "http://ngurajeka.com",
	}
	return c.RenderJson(page)
}
