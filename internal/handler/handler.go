package handler

import (
	"html/template"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB           *sqlx.DB
	HtmlTemplate *template.Template
}
