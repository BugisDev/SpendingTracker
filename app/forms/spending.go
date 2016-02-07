package forms

import (
	"github.com/revel/revel"
)

// AddSpendingTypeForm Handling Add New Category
type AddSpendingTypeForm struct {
	Data NewSpendingTypeData `json:"data"`
}

// NewCategoryData Part of AddCategoryForm
type NewSpendingTypeData struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

// Validate AddCategoryForm
func (f *AddSpendingTypeForm) Validate() (v revel.Validation) {

	// Validate Name
	v.Required(f.Data.Name).Key("Name")
	v.MinSize(f.Data.Name, 3).Key("Name")
	v.MaxSize(f.Data.Name, 32).Key("Name")

	// Validate Description
	v.MaxSize(f.Data.Description, 120).Key("Description")

	return v
}

// AddSpendingForm Handling Add New Spending
type AddSpendingForm struct {
	Data NewSpendingData `json:"data"`
}

// NewSpendingData Part of AddSpendingForm
type NewSpendingData struct {
	Name	string `json:"name"`
	Description string `json:"description"`
	Amount	float32 `json:"amount"`
	SpendingTypeID int `json:"spending_type_id"`
}

// Validate AddSpendingForm
func (f *AddSpendingForm) Validate() (v revel.Validation) {

	// Validate Name
	v.Required(f.Data.Name).Key("Name")
	v.MinSize(f.Data.Name, 2).Key("Name")
	v.MaxSize(f.Data.Name, 32).Key("Name")

	// Validate Description
	v.MaxSize(f.Data.Description, 120).Key("Description")

	// Validate Amount
	v.Required(f.Data.Amount).Key("Amount")

	// Validate TypeId
	v.Required(f.Data.SpendingTypeID).Key("SpendingTypeId")

	return v
}