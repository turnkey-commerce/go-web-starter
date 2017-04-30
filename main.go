package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/turnkey-commerce/go-web-starter/config"
	"github.com/turnkey-commerce/go-web-starter/controllers"
)

var (
	// Version is the Semver version of the application and should be passed in the build.
	version = "1.0"
	appName = "My Web App"
)

func main() {
	var err error
	// TODO: Setup logging
	startLog("Starting " + appName + " version " + version + ", website on port " +
		config.Settings.Website.HTTPPort + "...")
	// Validate the required config for the website.
	_, err = govalidator.ValidateStruct(config.Settings.Website)
	if err != nil {
		fatalError("Error in config.toml configuration:", err)
	}

	cookieKey := []byte(config.Settings.Website.CookieKey)
	secureCookie := config.Settings.Website.SecureHTTPS

	// Start the web server
	templates := controllers.PopulateTemplates("views")
	controllers.Register(templates, appName, version, cookieKey, secureCookie)
	err = http.ListenAndServe(":"+config.Settings.Website.HTTPPort, nil) // set listen port
	if err != nil {
		fatalError("ListenAndServe: ", err)
	}
}

func startLog(message string) {
	fmt.Println(message)
	log.Println(message)
}

func fatalError(message string, content interface{}) {
	fmt.Println(message, content)
	log.Fatal(message, content)
}
