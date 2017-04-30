package config_test

import (
	"testing"

	"github.com/turnkey-commerce/go-web-starter/config"
)

func TestWebsiteConfiguration(t *testing.T) {
	websiteSettings := config.Settings.Website

	if websiteSettings.HTTPPort != "8000" {
		t.Error("Config Website HTTPPort mismatch:\n", websiteSettings.HTTPPort)
	}

	if websiteSettings.CookieKey != "CookieEncryptionKey" {
		t.Error("Config Website HTTPPort mismatch:\n", websiteSettings.CookieKey)
	}

	if websiteSettings.SecureHTTPS != false {
		t.Error("Config Website SecureHTTPS mismatch:\n", websiteSettings.SecureHTTPS)
	}
}
