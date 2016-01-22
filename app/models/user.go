package models

import "time"

// User Model
type User struct {
	ID                int        `json:"id"`
	Username          string     `json:"username" sql:"not null;unique;type:varchar(64);unique_index"`
	Password          []byte     `json:"password" sql:"not null;unique"`
	Email             string     `json:"email" sql:"not null;unique;type:varchar(64);unique_index"`
	Fullname          string     `json:"full_name" sql:"not null;unique"`
	Profile           Profile    `json:"profile"`
	CreatedAt         time.Time  `json:"-"`
	UpdatedAt         time.Time  `json:"-"`
	DeletedAt         *time.Time `json:"-"`
	LastLoginAt       time.Time  `json:"-"`
	PasswordUpdatedAt time.Time  `json:"-"`
}

// Profile Model
// Related with User Model
type Profile struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId" sql:"index"`
	Gender      int8      `json:"gender"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
}
