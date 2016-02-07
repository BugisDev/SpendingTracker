package models

import (
	"time"

	"github.com/bugisdev/SpendingTracker/app/forms"
	"github.com/jinzhu/gorm"
	"github.com/bugisdev/helper"
	"github.com/bugisdev/usermodules"
)

// SpendingType Model
type SpendingType struct {
	ID          int        `json:"id"`
	Name        string     `json:"name" sql:"not null;type:varchar(64);"`
	Description string     `json:"description" sql:"not null;type:varchar(340);"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

// Spending Model
type Spending struct {
	ID             int              `json:"id"`
	UserID         usermodules.User `json:"userId"`
	Name           string           `json:"name" sql:"not null;type:varchar(64);"`
	Description    string           `json:"description" sql:"not null;type:varchar(340);"`
	SpendingType   SpendingType     `json:"spending_type"`
	SpendingTypeID int              `json:"-" sql:"index"`
	Amount         float32          `json:"amount" sql:"not null;double"`
	CreatedAt      time.Time        `json:"-"`
	UpdatedAt      time.Time        `json:"-"`
	DeletedAt      *time.Time       `json:"-"`
}

// NewSpendingType Function
func NewSpendingType(f forms.AddSpendingTypeForm, DB *gorm.DB) (spendingType SpendingType, err []helper.ErrorMessage) {

	v := f.Validate()

	if v.HasErrors() {
		for _, value := range v.Errors {
			err = append(err, helper.ErrorMessage{
				Code:    409,
				Source:  helper.SourceErrors{Pointer: value.Key},
				Title:   value.Message,
				Details: value.Message,
			})
		}

		return spendingType, err
	}

	spendingType.Name = f.Data.Name
	spendingType.Description = f.Data.Description

	_err := DB.Create(&spendingType).Error
	if _err != nil {
		err = append(err, helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{},
			Title:   "Failed Creating New Spending Type",
			Details: _err.Error(),
		})

		return spendingType, err
	}

	return spendingType, nil
}

// NewSpending Function
func NewSpending(f forms.AddSpendingForm, DB *gorm.DB) (spending Spending, err []helper.ErrorMessage) {

	v := f.Validate()

	if v.HasErrors() {
		for _, value := range v.Errors {
			err = append(err, helper.ErrorMessage{
				Code:    409,
				Source:  helper.SourceErrors{Pointer: value.Key},
				Title:   value.Message,
				Details: value.Message,
			})
		}

		return spending, err
	}

	spending.Name = f.Data.Name
	spending.Description = f.Data.Description
	spending.Amount = f.Data.Amount
	spending.SpendingTypeID = f.Data.SpendingTypeID

	_err := DB.Create(&spending).Error
	if _err != nil {
		err = append(err, helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{},
			Title:   "Failed Creating New Spending",
			Details: _err.Error(),
		})

		return spending, err
	}

	var spendingType SpendingType
	DB.Model(&spending).Related(&spendingType)
	spending.SpendingType = spendingType
	return spending, nil
}