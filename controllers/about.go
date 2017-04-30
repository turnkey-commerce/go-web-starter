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

func (controller *aboutController) get(rw http.ResponseWriter, req *http.Request) {
	var messages []string
	isAuthenticated := false
	userName := "test"
	vm := viewmodels.GetAboutViewModel(messages, controller.appName, isAuthenticated, userName, controller.version)
	controller.template.Execute(rw, vm)
}
