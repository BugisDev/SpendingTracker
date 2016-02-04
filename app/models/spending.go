package models

import (
	"time"

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
