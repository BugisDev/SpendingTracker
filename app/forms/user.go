package forms

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/bugisdev/SpendingTracker/app"
	"github.com/bugisdev/SpendingTracker/app/models"
	"github.com/oleiade/reflections"
)

// UserLoginForm Handling User Login
type UserLoginForm struct {
	Data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

// UserRegisterForm Handling User Registration
type UserRegisterForm struct {
	Data struct {
		Email           string `json:"email"`
		Fullname        string `json:"full_name"`
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
}

// Login Modules
func (login *UserLoginForm) Login() (models.User, []app.ErrorMessage) {
	var errorMessages []app.ErrorMessage
	var user models.User

	// Check Fields
	fields := []string{"Username", "Password"}
	for _, field := range fields {
		value, _ := reflections.GetField(login.Data, field)
		if value == "" {
			errorMessage := app.ErrorMessage{
				Code:    409,
				Source:  app.SourceError{Pointer: "/data/" + field},
				Title:   "Input Error",
				Details: "Field " + field + " are empty",
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	if len(errorMessages) > 0 {
		return user, errorMessages
	}

	err := app.DB.Where(&models.User{Username: login.Data.Username}).First(&user).Error
	if err != nil {
		errorMessage := app.ErrorMessage{
			Code:    400,
			Source:  app.SourceError{},
			Title:   err.Error(),
			Details: err.Error(),
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	if len(errorMessages) > 0 {
		return user, errorMessages
	}

	loggedIn := ComparePassword(&user, login.Data.Password)
	if loggedIn == false {
		errorMessage := app.ErrorMessage{
			Code:    409,
			Source:  app.SourceError{},
			Title:   "Wrong Username / Password Combination",
			Details: "Wrong Username / Password Combination",
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	return user, errorMessages
}

// ComparePassword Function
func ComparePassword(user *models.User, loginpassword string) bool {
	_loginpassword := sha256.Sum256([]byte(loginpassword))
	return subtle.ConstantTimeCompare(user.Password, _loginpassword[:]) == 1
}

// Register Modules
func (reg *UserRegisterForm) Register() (models.User, []app.ErrorMessage) {

	var errorMessages []app.ErrorMessage
	var user models.User

	// Check Fields
	fields := []string{"Email", "Fullname", "Username", "Password", "ConfirmPassword"}
	for _, field := range fields {
		value, _ := reflections.GetField(reg.Data, field)
		if value == "" {
			errorMessage := app.ErrorMessage{
				Code:    409,
				Source:  app.SourceError{Pointer: "/data/" + field},
				Title:   "Input Error",
				Details: "Field " + field + " are empty",
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	if len(errorMessages) > 0 {
		return user, errorMessages
	}

	// Check Username Length
	if len(reg.Data.Username) < 3 {
		errorMessage := app.ErrorMessage{
			Code:    409,
			Source:  app.SourceError{Pointer: "/data/username"},
			Title:   "Input Error",
			Details: "Username too short, Minimal 3 Character",
		}
		errorMessages = append(errorMessages, errorMessage)
	} else if len(reg.Data.Username) > 32 {
		errorMessage := app.ErrorMessage{
			Code:    409,
			Source:  app.SourceError{Pointer: "/data/username"},
			Title:   "Input Error",
			Details: "Username too long, Maximal 32 Character",
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	if len(errorMessages) > 0 {
		return user, errorMessages
	}

	// Check Password
	if reg.Data.Password != reg.Data.ConfirmPassword {
		errorMessage := app.ErrorMessage{
			Code:    409,
			Source:  app.SourceError{Pointer: "/data/password"},
			Title:   "Input Error",
			Details: "Password and Confirm Password Didn't Match",
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	if len(errorMessages) > 0 {
		return user, errorMessages
	}

	// save user to the database
	user.Username = reg.Data.Username
	user.Password = GeneratePassword(reg.Data.Password)
	user.Fullname = reg.Data.Fullname
	user.Email = reg.Data.Email
	err := app.DB.Create(&user).Error
	if err != nil {
		errorMessage := app.ErrorMessage{
			Code:    409,
			Source:  app.SourceError{},
			Title:   "Error When Registering User",
			Details: err.Error(),
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	return user, errorMessages
}

// GeneratePassword Function
func GeneratePassword(password string) []byte {
	_password := sha256.Sum256([]byte(password))
	return _password[:]
}
