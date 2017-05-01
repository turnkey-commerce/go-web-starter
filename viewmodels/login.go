package viewmodels

import (
	"html/template"
)

// LoginViewModel holds the view information for the login.gohtml template
type LoginViewModel struct {
	Title     string
	Nav       NavViewModel
	Messages  []string
	CsrfField template.HTML
}

// GetLoginViewModel populates the items required by the login.gohtml view
func GetLoginViewModel(messages []string, appName string) LoginViewModel {
	nav := NavViewModel{
		Active:          "login",
		IsAuthenticated: false,
		AppName:         appName,
	}

	result := LoginViewModel{
		Title:    appName + " - Login",
		Nav:      nav,
		Messages: messages,
	}
	return result
}
