package service

import (
	"html/template"
	"os"
	"strings"
)

// Parses encoded string to raw html with template.HTML
func ParseRawHtml(raw string) template.HTML {
    return template.HTML(raw)
}

// Return a split containing languages in APP_LANGUAGES env variable
func GetAppLanguages() []string {
    return strings.Split(os.Getenv("APP_LANGUAGES"), ",")
}

