package handler

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type GetPingResponse struct {
	Message string `json:"message"`
}

func (h *Handler) GetPing(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	_ = json.NewEncoder(ctx).Encode(&GetPingResponse{
		Message: "ok",
	})
}
