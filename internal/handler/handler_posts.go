package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kallydev/c10k/internal/database/model"
	"github.com/valyala/fasthttp"
)

type GetPostListErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) GetPostList(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	rows, err := h.DB.Queryx("SELECT * FROM posts;")
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)

		_ = json.NewEncoder(ctx).Encode(&GetPostListErrorResponse{
			Error: err.Error(),
		})

		return
	}

	var posts = make([]model.Post, 0)

	for rows.Next() {
		var post model.Post

		if err := rows.StructScan(&post); err != nil {
			ctx.SetStatusCode(http.StatusInternalServerError)

			_ = json.NewEncoder(ctx).Encode(&GetPostListErrorResponse{
				Error: err.Error(),
			})

			return
		}

		posts = append(posts, post)
	}

	_ = json.NewEncoder(ctx).Encode(posts)
}
