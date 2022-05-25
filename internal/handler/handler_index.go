package handler

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type GetIndexResponse struct {
	Message string `json:"message"`
}

func (h *Handler) GetIndex(ctx *fasthttp.RequestCtx) {
	_ = json.NewEncoder(ctx.Response.BodyWriter()).Encode(&GetIndexResponse{
		Message: "ok",
	})
}
