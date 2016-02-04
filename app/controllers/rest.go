package controllers

import (
	"github.com/bugisdev/helper"
	"github.com/revel/revel"
)

// RestController for RESTFul web service
type RestController struct {
	GormController
}

// Redirect Function to handle error
func (c RestController) Redirect(code int, errorMessages []helper.ErrorMessage) revel.Result {
	c.Response.WriteHeader(code, "application/json")
	return c.RenderJson(map[string]interface{}{
		"errors": errorMessages,
	})
}
