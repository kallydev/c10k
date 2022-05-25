package main

import (
	"log"
	"os"

	"github.com/fasthttp/router"
	"github.com/jmoiron/sqlx"
	"github.com/kallydev/c10k/internal/database"
	"github.com/kallydev/c10k/internal/handler"
	"github.com/valyala/fasthttp"
)

var db *sqlx.DB

func init() {
	var err error

	if db, err = database.Dial(os.Getenv("POSTGRES_DSN")); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	httpRouter := router.New()

	httpHandler := handler.Handler{
		DB: db,
	}

	httpRouter.GET("/", httpHandler.GetIndex)
	httpRouter.GET("/posts", httpHandler.GetPostList)

	if err := fasthttp.ListenAndServe(":80", httpRouter.Handler); err != nil {
		log.Fatalln(err)
	}
}
