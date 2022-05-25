package handler

import "github.com/jmoiron/sqlx"

type Handler struct {
	DB *sqlx.DB
}
