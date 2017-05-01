package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/turnkey-commerce/go-web-starter/viewmodels"
)

type loginController struct {
	template *template.Template
	appName  string
}

// get creates the "/login" form.
func (controller *loginController) get(w http.ResponseWriter, req *http.Request) {
	var messages []string
	vm := viewmodels.GetLoginViewModel(messages, controller.appName)
	vm.CsrfField = csrf.TemplateField(req)
	if err := controller.template.Execute(w, vm); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// post handles "/login" post requests.
func (controller *loginController) post(rw http.ResponseWriter, req *http.Request) {
	username := req.PostFormValue("username")
	// Do the processing of the user login
	//password := req.PostFormValue("password")
	fmt.Println(username)
}
