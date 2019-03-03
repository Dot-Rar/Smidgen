package handlers

import (
	"Smidgen/database"
	"github.com/valyala/fasthttp"
)

func NewHandler(ctx *fasthttp.RequestCtx) {
	content := string(ctx.QueryArgs().Peek("content"))

	if content == "" {
		ctx.Redirect("/", 400)
		return
	}

	if len(content) > 65535 {
		ctx.Redirect("/", 400)
		return
	}

	id := database.CreatePaste(content)

	ctx.Redirect("/" + id, 200)
}
