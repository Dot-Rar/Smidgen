package http

import (
	"Smidgen/config"
	"Smidgen/http/handlers"
	"Smidgen/http/ratelimit"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasttemplate"
	"io/ioutil"
	"log"
	"time"
)

func StartServer() {
	log.Println("Loading server")

	router := routing.New()

	// Handle static requests (assets)
	fs := &fasthttp.FS {
		Root: "./public/static",
		GenerateIndexPages: false,
		Compress: true,
	}
	staticHandler := fasthttp.CompressHandler(fs.NewRequestHandler())

	router.Get("/assets/*", func(ctx *routing.Context) error {
		staticHandler(ctx.RequestCtx)
		return nil
	})

	// Handle requests to /
	defaultHandler := fasthttp.CompressHandler(handlers.DefaultHandler)
	router.Get("/", func(ctx *routing.Context) error {
		defaultHandler(ctx.RequestCtx)
		return nil
	})

	// Handle new pastes
	lmt := tollbooth.NewLimiter(float64(config.Conf.Ratelimit.PastesPerHour / 60.0 / 60.0), &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Hour,
	})

	newHandler := ratelimit.LimitHandler(fasthttp.CompressHandler(handlers.NewHandler), lmt)
	router.Get("/new", func(ctx *routing.Context) error {
		newHandler(ctx.RequestCtx)
		return nil
	})

	// Handle requests to /raw
	rawHandler := fasthttp.CompressHandler(handlers.RawHandler)
	router.Get("/raw/<id>", func(ctx *routing.Context) error {
		ctx.RequestCtx.QueryArgs().Add("id", ctx.Param("id"))
		rawHandler(ctx.RequestCtx)
		return nil
	})

	// Handle paste requests
	pasteHandler:= fasthttp.CompressHandler(handlers.PasteHandler)
	router.Get("/<id>", func(ctx *routing.Context) error {
		ctx.RequestCtx.QueryArgs().Add("id", ctx.Param("id"))
		pasteHandler(ctx.RequestCtx)
		return nil
	})

	log.Println("Starting server")
	err := fasthttp.ListenAndServe(config.Conf.Server.Address, router.HandleRequest); if err != nil {
		panic(err)
	}
}

func LoadTemplates() {
	log.Println("Loading templates")

	index, err := ioutil.ReadFile("./public/templates/index.html"); if err != nil {
		panic(err)
	}
	handlers.DefaultTemplate = *fasttemplate.New(string(index), "{{", "}}")

	paste, err := ioutil.ReadFile("./public/templates/paste.html"); if err != nil {
		panic(err)
	}
	handlers.PasteTemplate = *fasttemplate.New(string(paste), "{{", "}}")
}
