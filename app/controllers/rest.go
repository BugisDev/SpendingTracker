package controllers

import "github.com/revel/revel"
import "github.com/bugisdev/SpendingTracker/app"

// RestController for RESTFul web service
type RestController struct {
	GormController
}

// Redirect Function to handle error
func (c RestController) Redirect(code int, errorMessages []app.ErrorMessage) revel.Result {
	c.Response.WriteHeader(code, "application/json")
	return c.RenderJson(map[string]interface{}{
		"errors": errorMessages,
	})
}
