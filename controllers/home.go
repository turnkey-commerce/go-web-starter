package controllers

import (
	"html/template"
	"net/http"

	"github.com/turnkey-commerce/go-web-starter/viewmodels"
)

type homeController struct {
	template *template.Template
	appName  string
}

func (controller *homeController) get(w http.ResponseWriter, r *http.Request) {
	var messages []string
	isAuthenticated := false
	userName := "test"
	vm := viewmodels.GetHomeViewModel(messages, controller.appName, isAuthenticated, userName)

	if err := controller.template.Execute(w, vm); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
