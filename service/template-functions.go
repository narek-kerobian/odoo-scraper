package service

import "html/template"

// Parses encoded string to raw html with template.HTML
func ParseRawHtml(raw string) template.HTML {
    return template.HTML(raw)
}
