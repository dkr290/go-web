package config

import (
	"html/template"
	"log"
)

// holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
