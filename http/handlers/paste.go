package handlers

import (
	"Smidgen/database"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
	"html"
)

var(
	PasteTemplate fasttemplate.Template

	/*style = styles.Get("swapoff")
	formatter = formatters.Get("html")*/
)

func PasteHandler(ctx *fasthttp.RequestCtx) {
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

	escaped := html.EscapeString(*content)
	/*var parsed string

	lexer := lexers.Analyse(*content)
	lexer := lexers.Get("html")
	fmt.Println(lexer == nil)
	if lexer == nil { // Not code
		parsed = escaped
	} else {
		iterator, err := lexer.Tokenise(nil, *content); if err != nil {
			parsed = escaped
		} else {
			buf := strings.Builder{}
			err := formatter.Format(&buf, style, iterator); if err != nil {
				parsed = escaped
			} else {
				parsed = buf.String()
			}
		}
	}*/

	res := PasteTemplate.ExecuteString(map[string]interface{}{
		"id": id,
		"content": escaped,
	})

	ctx.SetContentType("text/html; charset=utf8")
	ctx.SetStatusCode(200)

	_, err := fmt.Fprintln(ctx, res); if err != nil {
		ctx.SetStatusCode(500)
		fmt.Println(fmt.Sprintf("An error occurred while handling /%s: %s", id, err.Error()))
	}
}
