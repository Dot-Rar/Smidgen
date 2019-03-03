package handlers

import (
	"Smidgen/database"
	"fmt"
	"github.com/valyala/fasthttp"
)

func RawHandler(ctx *fasthttp.RequestCtx) {
	id := string(ctx.QueryArgs().Peek("id"))
	if id == "" {
		ctx.Redirect("/", 404)
		return
	}

	content := database.GetContent(id)
	if content == nil {
		ctx.Redirect("/", 404)
		return
	}

	ctx.SetContentType("text/plain; charset=utf8")
	ctx.SetStatusCode(200)

	_, err := fmt.Fprintln(ctx, *content); if err != nil {
		ctx.SetStatusCode(500)
		fmt.Println(fmt.Sprintf("An error occurred while handling /%s: %s", id, err.Error()))
	}
}
