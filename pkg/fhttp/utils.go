package fhttp

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func RenderJson(ctx *fasthttp.RequestCtx, statusCode int, res interface{}) {
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.SetStatusCode(statusCode)
	body, err := json.Marshal(res)
	if err != nil {
		ErrProcessRequest(ctx)
		return
	}
	if _, err = ctx.Write(body); err != nil {
		ErrProcessRequest(ctx)
		return
	}
}

func CombineHandlers(handler ...fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		for _, h := range handler {
			if ctx.Response.StatusCode() != 200 {
				return
			}
			h(ctx)
		}
	}
}
