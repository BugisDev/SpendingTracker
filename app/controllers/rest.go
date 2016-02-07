package controllers

import (
	"github.com/bugisdev/helper"
	"github.com/revel/revel"
	"strconv"
)

// RestController for RESTFul web service
type RestController struct {
	GormController
}

type Request struct {
	Limit, Number, Offset int
	Sort string
}

// Redirect Function to handle error
func (c RestController) Redirect(code int, errorMessages []helper.ErrorMessage) revel.Result {
	c.Response.WriteHeader(code, "application/json")
	return c.RenderJson(map[string]interface{}{
		"errors": errorMessages,
	})
}

// Parsing Request
func (c RestController) ParseRequest() (r Request) {

	q := c.Request.URL.Query()

	if q.Get("pageNumber") != "" {
		r.Number, _ = strconv.Atoi(q.Get("pageNumber"))
	} else {
		r.Number = 1
	}

	if q.Get("pageSize") != "" {
		r.Limit, _ = strconv.Atoi(q.Get("pageSize"))
	} else {
		r.Limit = 10
	}

	if q.Get("sort") != "" {
		r.Sort = q.Get("sort")
	} else {
		r.Sort = ""
	}

	if r.Number > 1 {
		r.Offset = (r.Number-1) * r.Limit
	} else {
		r.Offset = 0
	}

	return r
}