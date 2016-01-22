package controllers

import (
	"encoding/json"

	"github.com/bugisdev/SpendingTracker/app/forms"
	"github.com/revel/revel"
)

// SpendingController inherit RestController
// handling all Spending Resource
type SpendingController struct {
	RestController
}

// AddCategory Process
func (c SpendingController) AddCategory() revel.Result {
	var categoryForm forms.AddCategoryForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&categoryForm)

	category, err := categoryForm.AddCategory()

	if err != nil {
		return c.Redirect(400, err)
	}

	return c.RenderJson(category)
}

// AddSpending Process
func (c SpendingController) AddSpending() revel.Result {
	var spendingForm forms.AddSpendingForm
	body := json.NewDecoder(c.Request.Request.Body)
	body.Decode(&spendingForm)

	spending, err := spendingForm.AddSpending()

	if err != nil {
		return c.Redirect(400, err)
	}
	return c.RenderJson(spending)
}
