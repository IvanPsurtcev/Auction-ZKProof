package middleware

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"metamask-auth/pkg/jwt"
	"strings"
)

type (
	mwManger struct {
		logger      *zap.Logger
		jwtProvider jwt.HmacProvider
	}

	MwManager interface {
		JwtValidate() fasthttp.RequestHandler
	}
)

func NewMw(logger *zap.Logger) MwManager {
	return &mwManger{
		logger: logger,
	}
}

func (mw *mwManger) JwtValidate() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenRaw := strings.Split(string(ctx.Request.Header.Peek("Authorization")), " ")
		if len(tokenRaw) != 2 {
			ctx.Error("invalid token format", fasthttp.StatusBadRequest)
			return
		}
		claims, err := mw.jwtProvider.Verify(tokenRaw[1])
		if err != nil {
			ctx.Error("invalid token", fasthttp.StatusUnauthorized)
			return
		}
		ctx.SetUserValue("userAddress", claims.Subject)
	}
}
