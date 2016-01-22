package controllers

import (
	"encoding/json"

	"github.com/bugisdev/SpendingTracker/app/forms"
	"github.com/revel/revel"
)

// UserController inherit RestController
// handling all User Resource
type UserController struct {
	RestController
}

// Login Process
func (c UserController) Login() revel.Result {
	var login forms.UserLoginForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&login)

	user, err := login.Login()

	if err != nil {
		return c.Redirect(400, err)
	}

	c.Session["user_name"] = user.Username

	return c.RenderJson(user)
}

// Register Process
func (c UserController) Register() revel.Result {
	var reg forms.UserRegisterForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&reg)

	user, err := reg.Register()

	if err != nil {
		return c.Redirect(400, err)
	}

	c.Session["user_name"] = user.Username

	return c.RenderJson(user)
}
