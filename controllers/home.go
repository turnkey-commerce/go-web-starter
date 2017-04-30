package controllers

import (
	"html/template"
	"net/http"

	"github.com/turnkey-commerce/go-web-starter/viewmodels"
)

type homeController struct {
	template *template.Template
}

func (controller *homeController) get(rw http.ResponseWriter, req *http.Request) {
	var messages []string
	isAuthenticated := false
	userName := "test"
	vm := viewmodels.GetHomeViewModel(messages, isAuthenticated, userName)
	controller.template.Execute(rw, vm)
}
