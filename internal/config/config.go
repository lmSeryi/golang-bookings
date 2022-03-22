package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	Infolog       *log.Logger
	Errorlog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
