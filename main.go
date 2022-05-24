package main

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()

	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		_, _ = ctx.WriteString("hello world")
	})

	if err := fasthttp.ListenAndServe(":80", r.Handler); err != nil {
		log.Fatalln(err)
	}
}
