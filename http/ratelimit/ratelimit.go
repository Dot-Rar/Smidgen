package ratelimit

import (
	"Smidgen/config"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var(
	Pastes = make(map[string]int)
)

func RatelimitHandler(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		ip := ctx.RemoteIP().String()
		fmt.Println(ip)

		if _, contains := Pastes[ip]; contains {
			Pastes[ip] = Pastes[ip] + 1
		} else {
			Pastes[ip] = 1
		}

		go func() {
			time.Sleep(1 * time.Hour)

			Pastes[ip] = Pastes[ip] - 1
			if Pastes[ip] <= 0 {
				delete(Pastes, ip)
			}
		}()

		if Pastes[ip] <= config.Conf.Ratelimit.PastesPerHour { // We've added 1 already, so allow less than or equal to
			next(ctx)
		} else {
			ctx.Redirect("/", 429)
		}
	})
}