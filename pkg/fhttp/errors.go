package fhttp

import (
	"github.com/valyala/fasthttp"
)

func ErrProcessRequest(ctx *fasthttp.RequestCtx) {
	ctx.Error("can not process request", fasthttp.StatusInternalServerError)
}
