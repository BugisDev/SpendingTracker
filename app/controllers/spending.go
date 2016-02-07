package controllers

import (
	"encoding/json"

	"github.com/bugisdev/SpendingTracker/app/forms"
	"github.com/revel/revel"
	"github.com/bugisdev/SpendingTracker/app/models"
	"github.com/bugisdev/SpendingTracker/app"
)

// SpendingController inherit RestController
// handling all Spending Resource
type SpendingController struct {
	RestController
}

// AddType Process
func (c SpendingController) AddType() revel.Result {
	var f forms.AddSpendingTypeForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&f)

	spendingType, err := models.NewSpendingType(f, app.DB)

	if err != nil {
		return c.Redirect(400, err)
	}

	return c.RenderJson(spendingType)
}

// AddSpending Process
func (c SpendingController) AddSpending() revel.Result {
	var f forms.AddSpendingForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&f)

	spending, err := models.NewSpending(f, app.DB)

	if err != nil {
		return c.Redirect(400, err)
	}

	return c.RenderJson(spending)
}