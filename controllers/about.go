package controllers

import (
	"html/template"
	"net/http"

	"github.com/turnkey-commerce/go-web-starter/viewmodels"
)

type aboutController struct {
	template *template.Template
	version  string
	appName  string
}

func (controller *aboutController) get(w http.ResponseWriter, req *http.Request) {
	var messages []string
	isAuthenticated := false
	userName := "test"
	vm := viewmodels.GetAboutViewModel(messages, controller.appName, isAuthenticated, userName, controller.version)
	if err := controller.template.Execute(w, vm); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
