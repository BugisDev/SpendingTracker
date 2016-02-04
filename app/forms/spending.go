package forms

import (
	"github.com/bugisdev/SpendingTracker/app"
	"github.com/bugisdev/SpendingTracker/app/models"
	"github.com/bugisdev/helper"
	"github.com/oleiade/reflections"
)

// AddCategoryForm Handling Add New Category
type AddCategoryForm struct {
	Data struct {
		Name, Description string
	}
}

// AddSpendingForm Handling Add New Spending
type AddSpendingForm struct {
	Data struct {
		Name, Description string
		Amount            float32
		SpendingTypeID    int `json:"spending_type_id"`
	}
}

// AddCategory Modules
func (f *AddCategoryForm) AddCategory() (models.SpendingType, []helper.ErrorMessage) {
	var errorMessages []helper.ErrorMessage
	var category models.SpendingType

	// Check Fields
	fields := []string{"Name", "Description"}
	for _, field := range fields {
		value, _ := reflections.GetField(f.Data, field)
		if value == "" {
			errorMessage := helper.ErrorMessage{
				Code:    409,
				Source:  helper.SourceErrors{Pointer: "/data/" + field},
				Title:   "Input Error",
				Details: "Field " + field + " are empty",
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	if len(errorMessages) > 0 {
		return category, errorMessages
	}

	category.Name = f.Data.Name
	category.Description = f.Data.Description
	err := app.DB.Create(&category).Error
	if err != nil {
		errorMessage := helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{},
			Title:   "Failed Add New Category",
			Details: err.Error(),
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	return category, errorMessages
}

// AddSpending Modules
func (f *AddSpendingForm) AddSpending() (models.Spending, []helper.ErrorMessage) {
	var errorMessages []helper.ErrorMessage
	var spending models.Spending

	// Check Fields
	fields := []string{"Name"}
	for _, field := range fields {
		value, _ := reflections.GetField(f.Data, field)
		if value == "" {
			errorMessage := helper.ErrorMessage{
				Code:    409,
				Source:  helper.SourceErrors{Pointer: "/data/" + field},
				Title:   "Input Error",
				Details: "Field " + field + " are empty",
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	// Check Amount
	if f.Data.Amount == 0 {
		errorMessage := helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{Pointer: "/data/Amount"},
			Title:   "Input Error",
			Details: "Field Amount are 0",
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	// Check Spending Type Id
	if f.Data.SpendingTypeID == 0 {
		errorMessage := helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{Pointer: "/data/SpendingTypeID"},
			Title:   "Input Error",
			Details: "Field SpendingTypeID are empty",
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	if len(errorMessages) > 0 {
		return spending, errorMessages
	}

	spending.Name = f.Data.Name
	spending.Description = f.Data.Description
	spending.SpendingTypeID = f.Data.SpendingTypeID
	spending.Amount = f.Data.Amount
	err := app.DB.Create(&spending).Error
	if err != nil {
		errorMessage := helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{},
			Title:   "Failed Add New Spending",
			Details: err.Error(),
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	if len(errorMessages) > 0 {
		return spending, errorMessages
	}

	var spendingType models.SpendingType
	app.DB.Model(&spending).Related(&spendingType)
	spending.SpendingType = spendingType
	return spending, errorMessages
}

// GetAllSpendings User Records
func GetAllSpendings() ([]models.Spending, []helper.ErrorMessage) {

	var spending models.Spending
	var spendings []models.Spending
	var spendingtype models.SpendingType

	app.DB.Model(&spending).Find(&spendings)
	app.DB.Model(&spending).Related(&spendingtype, "SpendingType")

	return spendings, nil
}
