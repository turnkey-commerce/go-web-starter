package controllers

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

// Register the handlers for a given route.
func Register(templates *template.Template, appName string, version string, cookieKey []byte, secureCookie bool) {

	// setup CSRF protection for the post requests.
	CSRF := csrf.Protect(cookieKey, csrf.Secure(secureCookie))
	router := mux.NewRouter()

	hc := new(homeController)
	hc.template = templates.Lookup("home.gohtml")
	router.HandleFunc("/", hc.get)

	ac := new(aboutController)
	ac.template = templates.Lookup("about.gohtml")
	ac.version = version
	router.HandleFunc("/about", ac.get)

	lc := new(loginController)
	lc.template = templates.Lookup("login.gohtml")
	router.HandleFunc("/login", lc.get).Methods("GET")
	router.HandleFunc("/login", lc.post).Methods("POST")

	loc := new(logoutController)
	router.HandleFunc("/logout", loc.get)

	// Wrap the router in the CSRF protection.
	http.Handle("/", CSRF(router))

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/fonts/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".eot") {
		contentType = "application/vnd.ms-fontobject"
	} else if strings.HasSuffix(path, ".ttf") {
		contentType = "application/font-sfnt"
	} else if strings.HasSuffix(path, ".otf") {
		contentType = "application/font-sfnt"
	} else if strings.HasSuffix(path, ".woff") {
		contentType = "application/font-woff"
	} else if strings.HasSuffix(path, ".woff2") {
		contentType = "application/font-woff2"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

type appHandler func(http.ResponseWriter, *http.Request) (int, error)

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := fn(w, r); err != nil {
		switch status {
		// We can have cases as granular as we like, if we wanted to
		// return custom errors for specific status codes.
		case http.StatusInternalServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			// Catch any other errors we haven't explicitly handled
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// displayBool allows for colored display of Y/N in the templates.
func displayBool(input bool) template.HTML {
	if input {
		return "<span class=\"text-success\">Y<span>"
	}
	return "<span class=\"text-danger\">N<span>"
}

// displayActiveRow allows for light display of inactive rows in tables.
func displayActiveClass(input bool) template.HTMLAttr {
	if input {
		return "class=\"text-active\""
	}
	return "class=\"text-inactive\""
}

// PopulateTemplates loads and parses all of the templates in the templates directory
func PopulateTemplates(templatePath string) *template.Template {
	result := template.New("templates")

	basePath := templatePath
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
		}
	}

	var funcMap = template.FuncMap{
		"displayBool":        displayBool,
		"displayActiveClass": displayActiveClass,
	}

	result.Funcs(funcMap).ParseFiles(*templatePaths...)
	return result
}
