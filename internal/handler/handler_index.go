package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kallydev/c10k/internal/database/model"
	"github.com/valyala/fasthttp"
)

type GetIndexResponse struct {
	Message string `json:"message"`
}

type GetIndexTemplateData struct {
	Posts []model.Post
}

func (h *Handler) GetIndex(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	rows, err := h.DB.Queryx("SELECT * FROM posts;")
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)

		_ = json.NewEncoder(ctx.Response.BodyWriter()).Encode(&GetPostListErrorResponse{
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

	ctx.SetContentType("text/html")

	_ = h.HtmlTemplate.ExecuteTemplate(ctx, "index", &GetIndexTemplateData{
		Posts: posts,
	})
}
