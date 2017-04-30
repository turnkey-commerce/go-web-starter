package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/turnkey-commerce/go-ping-sites/viewmodels"
)

type loginController struct {
	template *template.Template
	appName  string
}

// get creates the "/login" form.
func (controller *loginController) get(rw http.ResponseWriter, req *http.Request) {
	var messages []string
	vm := viewmodels.GetLoginViewModel(messages)
	vm.CsrfField = csrf.TemplateField(req)
	controller.template.Execute(rw, vm)
}

// post handles "/login" post requests.
func (controller *loginController) post(rw http.ResponseWriter, req *http.Request) {
	username := req.PostFormValue("username")
	password := req.PostFormValue("password")
	fmt.Println(username, password)
}
