package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/bugisdev/helper"
	"github.com/bugisdev/usermodules"
	"github.com/ngurajeka/ngurajeka.com/app"
	"github.com/revel/revel"
)

// UserController inherit RestController
// handling all User Resource
type UserController struct {
	RestController
}

// Login Process
func (c UserController) Login() revel.Result {
	var f usermodules.UserLoginForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&f)

	user, err := usermodules.Login(f, app.DB)

	if err != nil {
		return c.Redirect(400, err)
	}

	c.Session["user_name"] = user.Username
	c.Session["user_id"] = strconv.Itoa(user.ID)

	return c.RenderJson(user)
}

// Register Process
func (c UserController) Register() revel.Result {
	var f usermodules.UserRegisterForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&f)

	user, err := usermodules.NewUser(f, app.DB)

	if err != nil {
		return c.Redirect(400, err)
	}

	c.Session["user_name"] = user.Username
	c.Session["user_id"] = strconv.Itoa(user.ID)

	return c.RenderJson(user)
}

// Update Process
func (c UserController) Update(id int) revel.Result {

	var f usermodules.UserUpdateForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&f)

	_id, _ := strconv.Atoi(c.Session["user_id"])
	if _id != id {
		var err []helper.ErrorMessage
		err = append(err, helper.ErrorMessage{
			Code:    400,
			Source:  helper.SourceErrors{},
			Title:   "Wrong UserID",
			Details: "",
		})

		return c.Redirect(400, err)
	}

	user, err := usermodules.UpdateSingle(id, f, app.DB)
	if err != nil {
		return c.Redirect(400, err)
	}

	return c.RenderJson(user)
}

// GetAll Process
func (c UserController) GetAll() revel.Result {

	r := c.ParseRequest()

	users, err := usermodules.GetAll(r.Limit, r.Offset, app.DB)
	if err != nil {
		return c.Redirect(400, err)
	}

	return c.RenderJson(users)
}