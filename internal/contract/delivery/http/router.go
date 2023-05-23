package contractHttp

import (
	"github.com/fasthttp/router"
	"metamask-auth/internal/contract"
	"metamask-auth/internal/middleware"
	"metamask-auth/pkg/fhttp"
)

func InitRouter(router *router.Group, h contract.Handlers, mw middleware.MwManager) {
	mwHandler := mw.JwtValidate()
	router.POST("/start", fhttp.CombineHandlers(mwHandler, h.StartAuction()))
	router.POST("/end", fhttp.CombineHandlers(mwHandler, h.EndAuction()))
	router.POST("/makeBid", fhttp.CombineHandlers(mwHandler, h.MakeBid()))
	router.POST("/bidWin", fhttp.CombineHandlers(mwHandler, h.BidWin()))
}
