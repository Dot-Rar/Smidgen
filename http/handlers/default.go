package handlers

import (
	"fmt"
	"github.com/valyala/fasthttp"
)
import "github.com/valyala/fasttemplate"

var(
	DefaultTemplate fasttemplate.Template
)

func DefaultHandler(ctx *fasthttp.RequestCtx) {
	res := DefaultTemplate.ExecuteString(map[string]interface{}{
	})

	ctx.SetContentType("text/html; charset=utf8")
	ctx.SetStatusCode(200)

	_, err := fmt.Fprintln(ctx, res); if err != nil {
		ctx.SetStatusCode(500)
		fmt.Println("An error occurred while handling / : " + err.Error())
	}
}
